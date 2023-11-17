### Example:

    package main
    
    import (
    "flag"
    "fmt"
    RIPEtools "github.com/OlegPowerC/RIPETools"
    )
    
    func main() {
    IpAddr := flag.String("i", "", "ip address, example: -i 8.8.8.8")
    flag.Parse()
    
        rp, rperr := RIPEtools.NewRIPEreq(*IpAddr)
        if rperr == nil {
            CNerr, CN := rp.GetCountry()
            if CNerr == nil {
                fmt.Println("Country:", CN)
            }
            NetErr, Net := rp.GetNetwork()
            if NetErr == nil {
                fmt.Println("Network:", Net)
            }
            MNTerr, MNT := rp.GetMaintainer()
            if MNTerr == nil {
                fmt.Println("Maintainer:", MNT)
            }
            ASerr, AS := rp.GetOriginAs()
            if ASerr == nil {
                fmt.Println("Origin AS", AS)
            }
        } else {
            fmt.Println(rperr)
        }
    }

#### result when use parameter -i 5.183.253.87

    Country: US
    Network: 5.183.253.0/24
    Maintainer: traffictransitsolution
    Origin AS 35830
