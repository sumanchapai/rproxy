Simple reverse proxy that sits infront of multiple services and choses a service
to serve request based on the request URL path.

Running this in launchctl as:

```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd"\>
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.user.accounting-reverse-proxy</string>

    <key>ProgramArguments</key>
    <array>
        <string>/opt/homebrew/bin/go</string>
        <string>run</string>
        <string>/Users/suman/Desktop/projects/rproxy/main.go</string>
    </array>

    <key>KeepAlive</key>
    <true/>

    <key>RunAtLoad</key>
    <true/>

    <key>StandardOutPath</key>
    <string>/tmp/accounting-reverse-proxy.log</string>

    <key>StandardErrorPath</key>
    <string>/tmp/accounting-reverse-proxy.err</string>
</dict>
</plist>

```
