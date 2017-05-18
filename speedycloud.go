package speedycloud

import (
    "fmt"
    "io/ioutil"
    "net"
    "strings"
    "time"

    "github.com/docker/machine/libmachine/drivers"
    "github.com/docker/machine/libmachine/log"
    "github.com/docker/machine/libmachine/mcnflag"
    "github.com/docker/machine/libmachine/mcnutils"
    "github.com/docker/machine/libmachine/ssh"
    "github.com/docker/machine/libmachine/state"
    "regexp"
)

type Driver struct {
    *drivers.BaseDriver
    ActiveTimeout    int
    SpeedCloudUrl    string
    ApiKey           string
    ApiSecret        string
    AvailabilityZone string
    MachineId        string
    KeyPairName      string
    NetworkName      string
    UserData         string
    PrivateKeyFile   string
    CpuNumber        int
    Memory           int
    DiskType         string
    DiskCapacity     int
    Isp              string
    Bandwidth        int
    ImageType        string
    IpType           string
    client           Client
}

const (
    defaultURL = "http://api.speedycloud.cn/api/v1/products"
    defaultSSHUser = "root"
    defaultSSHPort = 22
    defaultActiveTimeout = 200
    defaultKeyPairName = "cloudos"
    defaultAvailabilityZone = "SPC-BJ-15-A"
    defaultAPISecret = "26685451e2a302886fb2a53958a787f65d3433731bee56d8f7fe96a714f677db"
    defaultAPIKey = "A44384F4AE3998DFEEAFD5E6CC6B1D56"
    defaultKeyFile = "/tmp/cloudos"
    defaultNetworkName = ""
    defaultCpuNumber = 2
    defaultMemory = 1024
    defaultDiskType = "Normal"
    defaultDiskCapacity = 20
    defaultISP = "Private"
    defaultBandwidth = 2
    defaultImage = "Ubuntu 14.04"
    defaultIpType = "inner"
    defaultUserDate = "/tmp/userdata"
)

func (d *Driver) GetCreateFlags() []mcnflag.Flag {
    return []mcnflag.Flag{
        mcnflag.StringFlag{
            EnvVar: "SPEED_CLOUD_URL",
            Name:   "speedycloud-url",
            Usage:  "SpeedyCloud URL",
            Value:  defaultURL,
        },
        mcnflag.StringFlag{
            EnvVar: "SPEED_CLOUD_API_KEY",
            Name:   "speedycloud-api-key",
            Usage:  "SpeedyCloud api key",
            Value:  defaultAPIKey,
        },
        mcnflag.StringFlag{
            EnvVar: "SPEED_CLOUD_API_SECRET",
            Name:   "speedycloud-api-secret",
            Usage:  "SpeedyCloud password",
            Value:  defaultAPISecret,
        },
        mcnflag.StringFlag{
            EnvVar: "SPEED_CLOUD_KEYPAIR_NAME",
            Name:   "speedycloud-keypair-name",
            Usage:  "SpeedCloud keypair to use to SSH to the instance",
            Value:  defaultKeyPairName,
        },
        mcnflag.StringFlag{
            EnvVar: "SPEED_CLOUD_AVAILABILITY_ZONE",
            Name:   "speedycloud-availability-zone",
            Usage:  "SpeedyCloud zone where to create the instance ",
            Value:  defaultAvailabilityZone,
        },
        mcnflag.StringFlag{
            EnvVar: "SPEED_CLOUD_PRIVATE_KEY_FILE",
            Name:   "speedycloud-private-key-file",
            Usage:  "Private keyfile to use for SSH (absolute path)",
            Value:  defaultKeyFile,
        },
        mcnflag.StringFlag{
            EnvVar: "SPEED_CLOUD_USER_DATA_FILE",
            Name:   "speedycloud-user-data-file",
            Usage:  "File containing an SpeedyCloud userdata script",
            Value:  defaultUserDate,
        },
        mcnflag.StringFlag{
            EnvVar: "SPEED_CLOUD_NETWORK_NAME",
            Name:   "speedycloud-net-name",
            Usage:  "SpeedyCloud network name the machine will be connected on",
            Value:  defaultNetworkName,
        },

        mcnflag.StringFlag{
            EnvVar: "SPEED_CLOUD_SSH_USER",
            Name:   "speedycloud-ssh-user",
            Usage:  "SpeedCloud SSH user",
            Value:  defaultSSHUser,
        },
        mcnflag.IntFlag{
            EnvVar: "SPEED_CLOUD_SSH_PORT",
            Name:   "speedycloud-ssh-port",
            Usage:  "SpeedyCloud SSH port",
            Value:  defaultSSHPort,
        },
        mcnflag.IntFlag{
            EnvVar: "SPEED_CLOUD_ACTIVE_TIMEOUT",
            Name:   "speedycloud-active-timeout",
            Usage:  "SpeedyCloud active timeout",
            Value:  defaultActiveTimeout,
        },
        mcnflag.IntFlag{
            EnvVar: "SPEED_CLOUD_CPU_NUMBER",
            Name:   "speedycloud-cpu-number",
            Usage:  "cpu number of the instance",
            Value:  defaultCpuNumber,
        },
        mcnflag.IntFlag{
            EnvVar: "SPEED_CLOUD_MEMORY",
            Name:   "speedycloud-memory",
            Usage:  "memory of the instance",
            Value:  defaultMemory,
        },
        mcnflag.StringFlag{
            EnvVar: "SPEED_CLOUD_DISKTYPE",
            Name:   "speedycloud-disk-type",
            Usage:  "disk type of the instance",
            Value:  defaultDiskType,
        },
        mcnflag.IntFlag{
            EnvVar: "SPEED_CLOUD_DISKCAPACITY",
            Name:   "speedycloud-disk-capacity",
            Usage:  "disk capacity of the instance",
            Value:  defaultDiskCapacity,
        },
        mcnflag.StringFlag{
            EnvVar: "SPEED_CLOUD_ISP",
            Name:   "speedycloud-isp",
            Usage:  "isp type of the instance",
            Value:  defaultISP,
        },
        mcnflag.IntFlag{
            EnvVar: "SPEED_CLOUD_BANDWIDTH",
            Name:   "speedycloud-bandwidth",
            Usage:  "bandwidth type of the instance",
            Value:  defaultBandwidth,
        },
        mcnflag.StringFlag{
            EnvVar: "SPEED_CLOUD_IMAGE_TYPE",
            Name:   "speedycloud-image-type",
            Usage:  "imgage type of the instance",
            Value:  defaultImage,
        },
        mcnflag.StringFlag{
            EnvVar: "SPEED_CLOUD_IP_TYPE",
            Name:   "speedycloud-ip-type",
            Usage:  "ip type of the instance, cloud be outer or inner",
            Value:  defaultIpType,
        },
    }
}

func NewDriver(hostName, storePath string) drivers.Driver {
    return NewDerivedDriver(hostName, storePath)
}

func NewDerivedDriver(hostName, storePath string) *Driver {
    return &Driver{
        client:        &GenericClient{},
        ActiveTimeout: defaultActiveTimeout,
        BaseDriver: &drivers.BaseDriver{
            SSHUser:     defaultSSHUser,
            SSHPort:     defaultSSHPort,
            MachineName: hostName,
            StorePath:   storePath,
        },
    }
}

func (d *Driver) GetSSHHostname() (string, error) {
    return d.GetIP()
}

func (d *Driver) SetClient(client Client) {
    d.client = client
}

// DriverName returns the name of the driver
func (d *Driver) DriverName() string {
    return "speedycloud"
}

func (d *Driver) SetConfigFromFlags(flags drivers.DriverOptions) error {
    d.SpeedCloudUrl = flags.String("speedycloud-url")
    d.ApiKey = flags.String("speedycloud-api-key")
    d.ApiSecret = flags.String("speedycloud-api-secret")
    d.KeyPairName = flags.String("speedycloud-keypair-name")
    d.AvailabilityZone = flags.String("speedycloud-availability-zone")
    d.PrivateKeyFile = flags.String("speedycloud-private-key-file")
    d.NetworkName = flags.String("speedycloud-net-name")
    d.SSHUser = flags.String("speedycloud-ssh-user")
    d.SSHPort = flags.Int("speedycloud-ssh-port")
    d.ActiveTimeout = flags.Int("speedycloud-active-timeout")
    d.CpuNumber = flags.Int("speedycloud-cpu-number")
    d.Memory = flags.Int("speedycloud-memory")
    d.DiskType = flags.String("speedycloud-disk-type")
    d.DiskCapacity = flags.Int("speedycloud-disk-capacity")
    d.Isp = flags.String("speedycloud-isp")
    d.Bandwidth = flags.Int("speedycloud-bandwidth")
    d.ImageType = flags.String("speedycloud-image-type")
    d.IpType = flags.String("speedycloud-ip-type")
    if flags.String("speedycloud-user-data-file") != "" {
        userData, err := ioutil.ReadFile(flags.String("speedycloud-user-data-file"))
        if err == nil {
            d.UserData = string(userData[:])
        } else {
            return err
        }
    }

    d.SetSwarmConfigFromFlags(flags)

    return d.checkConfig()
}

func (d *Driver) GetURL() (string, error) {
    if err := drivers.MustBeRunning(d); err != nil {
        return "", err
    }

    ip, err := d.GetIP()
    if err != nil {
        return "", err
    }
    if ip == "" {
        return "", nil
    }

    return fmt.Sprintf("tcp://%s", net.JoinHostPort(ip, "2376")), nil
}

func (d *Driver) GetIP() (string, error) {
    if d.IPAddress != "" {
        return d.IPAddress, nil
    }

    log.Debug("Looking for the IP address...", map[string]string{"MachineId": d.MachineId})

    if err := d.initCompute(); err != nil {
        return "", err
    }

    // Looking for the IP address in a retry loop to deal with OpenStack latency
    for retryCount := 0; retryCount < 200; retryCount++ {
        addresses, err := d.client.GetInstanceIPAddresses(d)
        if err != nil {
            return "", err
        }
        if d.IpType == "inner" {
            for _, a := range addresses {
                if match, _:= regexp.MatchString(`192.*|172.*|10.*`, a); match{
                    log.Debug("the address is %s", a)
                    return a, nil
                }
            }
        } else {
            for _, a := range addresses {
                if match, _:= regexp.MatchString(`192.*|172.*|10.*`, a); !match{
                    log.Info("the address is %s", a)
                    return a, nil
                }

            }
        }

        time.Sleep(2 * time.Second)

    }
    return "", fmt.Errorf("No IP found for the machine")
}

func (d *Driver) GetState() (state.State, error) {
    log.Debug("Get status for OpenStack instance...", map[string]string{"MachineId": d.MachineId})
    if err := d.initCompute(); err != nil {
        return state.None, err
    }

    s, err := d.client.GetInstanceState(d)
    if err != nil {
        return state.None, err
    }

    log.Debug("State for OpenStack instance", map[string]string{
        "MachineId": d.MachineId,
        "State":     s,
    })

    switch s {
    case "Running":
        return state.Running, nil
    case "Suspended":
        return state.Saved, nil
    case "Stopped":
        return state.Stopped, nil
    case "Provisioning":
        return state.Starting, nil
    case "ERROR":
        return state.Error, nil
    }
    return state.None, nil
}

func (d *Driver) Create() error {
    if err := d.resolveIds(); err != nil {
        return err
    }
    if d.KeyPairName != "" {
        if err := d.loadSSHKey(); err != nil {
            return err
        }
    } else {
        d.KeyPairName = fmt.Sprintf("%s-%s", d.MachineName, mcnutils.GenerateRandomID())
        if err := d.createSSHKey(); err != nil {
            return err
        }
    }
    if err := d.createMachine(); err != nil {
        return err
    }
    if err := d.waitForInstanceActive(); err != nil {
        return err
    }
    if err := d.lookForIPAddress(); err != nil {
        return err
    }
    return nil
}

func (d *Driver) Start() error {
    if err := d.initCompute(); err != nil {
        return err
    }

    return d.client.StartInstance(d)
}

func (d *Driver) Stop() error {
    if err := d.initCompute(); err != nil {
        return err
    }

    return d.client.StopInstance(d)
}

func (d *Driver) Restart() error {
    if err := d.initCompute(); err != nil {
        return err
    }

    return d.client.RestartInstance(d)
}

func (d *Driver) Kill() error {

    return d.Stop()
}

func (d *Driver) Remove() error {
    log.Debug("deleting instance...", map[string]string{"MachineId": d.MachineId})
    log.Info("Deleting speedycloud instance...")
    if err := d.initCompute(); err != nil {
        return err
    }
    if status, _ := d.GetState(); status != state.Stopped {
        d.Stop()
        d.waitForInstanceStopped()
    }
    if err := d.client.DeleteInstance(d); err != nil {
        return err
    }
    //log.Debug("deleting key pair...", map[string]string{"Name": d.KeyPairName})
    // TODO (fsoppelsa) maybe we want to check this, in case of shared keypairs, before removal
    //if err := d.client.DeleteKeyPair(d, d.KeyPairName); err != nil {
    //	return err
    //}
    return nil
}

const (
    errorMandatoryEnvOrOption string = "%s must be specified either using the environment variable %s or the CLI option %s"
    //errorMandatoryOption string = "%s must be specified using the CLI option %s"
    //errorExclusiveOptions string = "Either %s or %s must be specified, not both"
    //errorBothOptions string = "Both %s and %s must be specified"
    //errorMandatoryTenantNameOrID string = "Tenant id or name must be provided either using one of the environment variables OS_TENANT_ID and OS_TENANT_NAME or one of the CLI options --speedycloud-tenant-id and --speedycloud-tenant-name"
    //errorWrongEndpointType string = "Endpoint type must be 'publicURL', 'adminURL' or 'internalURL'"
    //errorUnknownFlavorName string = "Unable to find flavor named %s"
    //errorUnknownImageName string = "Unable to find image named %s"
    //errorUnknownNetworkName string = "Unable to find network named %s"
    //errorUnknownTenantName string = "Unable to find tenant named %s"
)

func (d *Driver) checkConfig() error {

    if d.SpeedCloudUrl == "" {
        return fmt.Errorf(errorMandatoryEnvOrOption, "Speedycloud url", "SPEED_CLOUD_URL", "--speedycloud-url")
    }

    if d.ApiKey == "" {
        return fmt.Errorf(errorMandatoryEnvOrOption, "Api key", "SPEED_CLOUD_API_KEY", "speedycloud-api-key")
    }
    if d.ApiSecret == "" {
        return fmt.Errorf(errorMandatoryEnvOrOption, "Api secret", "SPEED_CLOUD_API_SECRET", "speedycloud-api-secret")
    }
    //if d.KeyPairName == "" {
    //    return fmt.Errorf(errorMandatoryEnvOrOption, "Key Pair Name", "SPEED_CLOUD_KEYPAIR_NAME", "speedycloud-keypair-name")
    //}
    if d.AvailabilityZone == "" {
        return fmt.Errorf(errorMandatoryEnvOrOption, "AZ", "SPEED_CLOUD_AVAILABILITY_ZONE", "speedycloud-availability-zone")
    }
    //if d.PrivateKeyFile == "" {
    //    return fmt.Errorf(errorMandatoryEnvOrOption, "Private Key File", "SPEED_CLOUD_PRIVATE_KEY_FILE", "speedycloud-private-key-file")
    //}
    if d.SSHUser == "" {
        return fmt.Errorf(errorMandatoryEnvOrOption, "Ssh User ", "SPEED_CLOUD_SSH_USER", "speedycloud-ssh-user")
    }
    if string(d.SSHPort) == "" {
        return fmt.Errorf(errorMandatoryEnvOrOption, "Ssh Port", "SPEED_CLOUD_SSH_PORT", "speedycloud-ssh-port")
    }
    if string(d.CpuNumber) == "" {
        return fmt.Errorf(errorMandatoryEnvOrOption, "Cpu Number", "SPEED_CLOUD_CPU_NUMBER", "speedycloud-cpu-number")
    }
    if string(d.Memory) == "" {
        return fmt.Errorf(errorMandatoryEnvOrOption, "Memory", "SPEED_CLOUD_MEMORY", "speedycloud-memory")
    }
    if d.ImageType == "" {
        return fmt.Errorf(errorMandatoryEnvOrOption, "Key Pair Name", "SPEED_CLOUD_IMAGE_TYPE", "speedycloud-image-type")
    }

    return nil
}

func (d *Driver) resolveIds() error {

    return nil
}

func (d *Driver) initCompute() error {
    if err := d.client.InitProviderClient(d); err != nil {
        return err
    }
    if err := d.client.InitComputeClient(d); err != nil {
        return err
    }
    return nil
}

func (d *Driver) initNetwork() error {
    if err := d.client.InitProviderClient(d); err != nil {
        return err
    }
    if err := d.client.InitNetworkClient(d); err != nil {
        return err
    }
    return nil
}

func (d *Driver) loadSSHKey() error {
    log.Debug("Loading Key Pair", d.KeyPairName)
    if err := d.initCompute(); err != nil {
        return err
    }
    log.Debug("Loading Private Key from", d.PrivateKeyFile)
    privateKey, err := ioutil.ReadFile(d.PrivateKeyFile)
    if err != nil {
        return err
    }
    publicKey, err := d.client.GetPublicKey(d.KeyPairName)
    if err != nil {
        return err
    }
    if err := ioutil.WriteFile(d.privateSSHKeyPath(), privateKey, 0600); err != nil {
        return err
    }
    if err := ioutil.WriteFile(d.publicSSHKeyPath(), publicKey, 0600); err != nil {
        return err
    }

    return nil
}

func (d *Driver) createSSHKey() error {
    sanitizeKeyPairName(&d.KeyPairName)
    log.Debug("Creating Key Pair...", map[string]string{"Name": d.KeyPairName})
    if err := ssh.GenerateSSHKey(d.GetSSHKeyPath()); err != nil {
        return err
    }
    publicKey, err := ioutil.ReadFile(d.publicSSHKeyPath())
    if err != nil {
        return err
    }

    if err := d.initCompute(); err != nil {
        return err
    }
    if err := d.client.CreateKeyPair(d, d.KeyPairName, string(publicKey)); err != nil {
        return err
    }
    return nil
}

func (d *Driver) createMachine() error {
    log.Debug("Creating SpeedyCloud instance...", map[string]string{
        "Cpu": string(d.CpuNumber),
        "Memory": string(d.Memory),
        "diskCapacity": string(d.DiskCapacity),
        "ImageId":  d.ImageType,
    })

    if err := d.initCompute(); err != nil {
        return err
    }
    instanceID, err := d.client.CreateInstance(d)
    if err != nil {
        return err
    }
    d.MachineId = instanceID
    return nil
}

//func (d *Driver) assignFloatingIP() error {
//    var err error
//
//    if d.ComputeNetwork {
//        err = d.initCompute()
//    } else {
//        err = d.initNetwork()
//    }
//
//    if err != nil {
//        return err
//    }
//
//    ips, err := d.client.GetFloatingIPs(d)
//    if err != nil {
//        return err
//    }
//
//    var floatingIP *FloatingIP
//
//    log.Debugf("Looking for an available floating IP", map[string]string{
//        "MachineId": d.MachineId,
//        "Pool":      d.FloatingIpPool,
//    })
//
//    for _, ip := range ips {
//        if ip.PortId == "" {
//            log.Debug("Available floating IP found", map[string]string{
//                "MachineId": d.MachineId,
//                "IP":        ip.Ip,
//            })
//            floatingIP = &ip
//            break
//        }
//    }
//
//    if floatingIP == nil {
//        floatingIP = &FloatingIP{}
//        log.Debug("No available floating IP found. Allocating a new one...", map[string]string{"MachineId": d.MachineId})
//    } else {
//        log.Debug("Assigning floating IP to the instance", map[string]string{"MachineId": d.MachineId})
//    }
//
//    if err := d.client.AssignFloatingIP(d, floatingIP); err != nil {
//        return err
//    }
//    d.IPAddress = floatingIP.Ip
//    return nil
//}

func (d *Driver) waitForInstanceActive() error {
    log.Debug("Waiting for the SpeedyCloud instance to be running...", map[string]string{"MachineId": d.MachineId})
    if err := d.client.WaitForInstanceStatus(d, "Running"); err != nil {
        return err
    }
    return nil
}

func (d *Driver) waitForInstanceStopped() error {
    log.Debug("Waiting for the SpeedyCloud instance to be stopped...", map[string]string{"MachineId": d.MachineId})
    if err := d.client.WaitForInstanceStatus(d, "Stopped"); err != nil {
        return err
    }
    return nil
}

func (d *Driver) lookForIPAddress() error {
    ip, err := d.GetIP()
    if err != nil {
        return err
    }
    d.IPAddress = ip
    log.Debug("IP address found", map[string]string{
        "IP":        ip,
        "MachineId": d.MachineId,
    })
    return nil
}

func (d *Driver) privateSSHKeyPath() string {
    return d.GetSSHKeyPath()
}

func (d *Driver) publicSSHKeyPath() string {
    return d.GetSSHKeyPath() + ".pub"
}

func sanitizeKeyPairName(s *string) {
    *s = strings.Replace(*s, ".", "_", -1)
}
