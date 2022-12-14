package main

import (
    "fmt"
    "log"
    "net/http"
    "os/exec"
    
    "github.com/gorilla/mux"
)

type BastileResponse struct {
    Target string
    Release string
}

func BastilleStart(response http.ResponseWriter, router *http.Request) {
    // bastille start TARGET
    vars := mux.Vars(router)
    target := vars["target"]
    out, err := exec.Command("bastille", "start", target).Output()
    
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Executing: bastille start %s\n", target)
    fmt.Println(string(out))

    response.WriteHeader(http.StatusOK)
    response.Write([]byte(out))
}

func BastilleStop(response http.ResponseWriter, router *http.Request) {
    // bastille stop TARGET
    vars := mux.Vars(router)
    target := vars["target"]
    out, err := exec.Command("bastille", "stop", target).Output()
    
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Executing: bastille stop %s\n", target)
    fmt.Println(string(out))

    response.WriteHeader(http.StatusOK)
    response.Write([]byte(out))
}

func BastilleRestart(response http.ResponseWriter, router *http.Request) {
    // bastille restart TARGET
    vars := mux.Vars(router)
    target := vars["target"]
    out, err := exec.Command("bastille", "restart", target).Output()
    
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Executing: bastille restart %s\n", target)
    fmt.Println(string(out))

    response.WriteHeader(http.StatusOK)
    response.Write([]byte(out))
}

func BastillePkg(response http.ResponseWriter, router *http.Request) {
    // bastille pkg TARGET args
    vars := mux.Vars(router)
    target := vars["target"]
    args := vars["args"]
    out, err := exec.Command("bastille", "pkg", target, args).Output()
    
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Executing: bastille pkg %s %s\n", target, args)
    fmt.Println(string(out))

    response.WriteHeader(http.StatusOK)
    response.Write([]byte(out))
}

func BastillePkgInstall(response http.ResponseWriter, router *http.Request) {
    // bastille pkg TARGET install args -y
    vars := mux.Vars(router)
    target := vars["target"]
    args := vars["args"]
    out, err := exec.Command("bastille", "pkg", target, "install", args, "-y").Output()
    
    if err != nil {
        fmt.Printf("DEBUG: error encountered")
    }

    fmt.Printf("Executing: bastille pkg %s install %s -y\n", target, args)
    fmt.Println(string(out))

    response.WriteHeader(http.StatusOK)
    response.Write([]byte(out))
}

func BastillePkgUpgrade(response http.ResponseWriter, router *http.Request) {
    // bastille pkg TARGET upgrade -y
    vars := mux.Vars(router)
    target := vars["target"]
    out, err := exec.Command("bastille", "pkg", target, "upgrade", "-y").Output()
    
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Executing: bastille pkg %s upgrade -y\n", target)
    fmt.Println(string(out))

    response.WriteHeader(http.StatusOK)
    response.Write([]byte(out))
}

func BastilleList(response http.ResponseWriter, router *http.Request) {
    // bastille list
    out, err := exec.Command("bastille", "list", "-a").Output()
    
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Executing: bastille list -a\n")
    fmt.Println(string(out))

    response.WriteHeader(http.StatusOK)
    response.Write([]byte(out))
}

func BastilleUpdate(response http.ResponseWriter, router *http.Request) {
    // bastille update RELEASE
    vars := mux.Vars(router)
    release := vars["release"]
    out, err := exec.Command("bastille", "update", release).Output()
    
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Executing: bastille update %q\n", release)
    fmt.Println(string(out))

    response.WriteHeader(http.StatusOK)
    response.Write([]byte(out))
}

func BastilleCmd(response http.ResponseWriter, router *http.Request) {
    // bastille cmd TARGET args
    vars := mux.Vars(router)
    target := vars["target"]
    args := vars["args"]
    out, err := exec.Command("bastille", "cmd", target, args).Output()
    
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Executing: bastille cmd %s %s\n", target, args)
    fmt.Println(string(out))

    response.WriteHeader(http.StatusOK)
    response.Write([]byte(out))
}

func BastilleService(response http.ResponseWriter, router *http.Request) {
    // bastille service TARGET args
    vars := mux.Vars(router)
    target := vars["target"]
    service := vars["service"]
    args := vars["args"]
    out, err := exec.Command("bastille", "service", target, service, args).Output()
    
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Executing: bastille service %s %s %s\n", target, service, args)
    fmt.Println(string(out))

    response.WriteHeader(http.StatusOK)
    response.Write([]byte(out))
}

func BastilleSysrc(response http.ResponseWriter, router *http.Request) {
    // bastille sysrc TARGET args
    vars := mux.Vars(router)
    target := vars["target"]
    args := vars["args"]
    out, err := exec.Command("bastille", "sysrc", target, args).Output()
    
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Executing: bastille sysrc %s %s\n", target, args)
    fmt.Println(string(out))

    response.WriteHeader(http.StatusOK)
    response.Write([]byte(out))
}

func BastilleBootstrap(response http.ResponseWriter, router *http.Request) {
    // bastille bootstrap RELEASE
    vars := mux.Vars(router)
    release := vars["release"]
    out, err := exec.Command("bastille", "bootstrap", release).Output()
    
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Executing: bastille bootstrap %s\n", release)
    fmt.Println(string(out))

    response.WriteHeader(http.StatusOK)
    response.Write([]byte(out))
}

func BastilleConfig(response http.ResponseWriter, router *http.Request) {
    // bastille config TARGET args
    vars := mux.Vars(router)
    target := vars["target"]
    args := vars["args"]
    out, err := exec.Command("bastille", "config", target, args).Output()
    
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Executing: bastille config %s %s\n", target, args)
    fmt.Println(string(out))

    response.WriteHeader(http.StatusOK)
    response.Write([]byte(out))
}

func BastilleClone(response http.ResponseWriter, router *http.Request) {
    // bastille clone TARGET newname IP
    vars := mux.Vars(router)
    target := vars["target"]
    newname := vars["newname"]
    IP := vars["IP"]
    out, err := exec.Command("bastille", "clone", target, newname, IP).Output()
    
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Executing: bastille clone %s %s %s\n", target, newname, IP)
    fmt.Println(string(out))

    response.WriteHeader(http.StatusOK)
    response.Write([]byte(out))
}

func BastilleCreate(response http.ResponseWriter, router *http.Request) {
    // bastille create name release IP
    vars := mux.Vars(router)
    name := vars["name"]
    release := vars["release"]
    IP := vars["IP"]
    out, err := exec.Command("bastille", "create", name, release, IP).Output()
    
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Executing: bastille create %s %s %s\n", name, release, IP)
    fmt.Println(string(out))

    response.WriteHeader(http.StatusOK)
    response.Write([]byte(out))
}

func BastilleDestroy(response http.ResponseWriter, router *http.Request) {
    // bastille destroy TARGET
    vars := mux.Vars(router)
    target := vars["target"]
    out, err := exec.Command("bastille", "destroy", target).Output()
    
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Executing: bastille destroy %s\n", target)
    fmt.Println(string(out))

    response.WriteHeader(http.StatusOK)
    response.Write([]byte(out))
}

func BastilleRdr(response http.ResponseWriter, router *http.Request) {
    // bastille rdr TARGET proto hostport jailport
    vars := mux.Vars(router)
    target := vars["target"]
    proto := vars["proto"]
    hostport := vars["hostport"]
    jailport := vars["jailport"]
    out, err := exec.Command("bastille", "rdr", target, proto, hostport, jailport).Output()
    
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Executing: bastille rdr %s %s %d %d\n", target, proto, hostport, jailport)
    fmt.Println(string(out))

    response.WriteHeader(http.StatusOK)
    response.Write([]byte(out))
}

func BastilleTemplate(response http.ResponseWriter, router *http.Request) {
    // bastille template TARGET project/repo
    vars := mux.Vars(router)
    target := vars["target"]
    project := vars["project"]
    repo := vars["repo"]
    template := project + "/" + repo
    out, err := exec.Command("bastille", "template", target, template).Output()
    
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Executing: bastille template %s %s\n", target, template)
    fmt.Println(string(out))

    response.WriteHeader(http.StatusOK)
    response.Write([]byte(out))
}

func main() {
    fmt.Println("Starting the Bastille API...")
    
    router := mux.NewRouter()
    router.HandleFunc("/v1/bastille/list", BastilleList).Methods("GET")
    router.HandleFunc("/v1/bastille/bootstrap/{release}", BastilleBootstrap).Methods("POST")
    router.HandleFunc("/v1/bastille/clone/{target}/{newname}/{IP}", BastilleClone).Methods("POST")
    router.HandleFunc("/v1/bastille/cmd/{target}/{args}", BastilleCmd).Methods("POST")
    router.HandleFunc("/v1/bastille/config/{target}/{args}", BastilleConfig).Methods("POST")
    router.HandleFunc("/v1/bastille/create/{name}/{release}/{IP}", BastilleCreate).Methods("POST")
    router.HandleFunc("/v1/bastille/destroy/{target}", BastilleDestroy).Methods("POST")
    router.HandleFunc("/v1/bastille/pkg/{target}/install/{args}", BastillePkgInstall).Methods("POST")
    router.HandleFunc("/v1/bastille/pkg/{target}/upgrade", BastillePkgUpgrade).Methods("POST")
    router.HandleFunc("/v1/bastille/pkg/{target}/{args}", BastillePkg).Methods("POST")
    router.HandleFunc("/v1/bastille/rdr/{target}/{proto}/{hostport}/{jailport}", BastilleRdr).Methods("POST")
    router.HandleFunc("/v1/bastille/restart/{target}", BastilleRestart).Methods("POST")
    router.HandleFunc("/v1/bastille/service/{target}/{service}/{args}", BastilleService).Methods("POST")
    router.HandleFunc("/v1/bastille/start/{target}", BastilleStart).Methods("POST")
    router.HandleFunc("/v1/bastille/stop/{target}", BastilleStop).Methods("POST")
    router.HandleFunc("/v1/bastille/sysrc/{target}/{args}", BastilleSysrc).Methods("POST")
    router.HandleFunc("/v1/bastille/template/{target}/{project}/{repo}", BastilleTemplate).Methods("POST")
    router.HandleFunc("/v1/bastille/update/{release}", BastilleUpdate).Methods("POST")
    fmt.Println("Listening on port 12345")
    http.ListenAndServe(":12345", router)
}
