# Setting Up Leither Autostart Service on macOS

This guide explains how to configure Leither to start automatically when you log in to macOS.

## Overview

macOS uses `launchd` with plist files to manage services. This guide sets up Leither as a LaunchAgent that runs when you log in.

## Prerequisites

- Leither binary at `/Users/cfa532/Documents/GitHub/darwin/Leither`
- macOS with user account
- **Full Disk Access** permission granted to Terminal (or your preferred terminal app)

## Step 1: Grant Full Disk Access Permission

**Important**: Leither requires Full Disk Access to run properly as a service on macOS. This permission is required for Leither to access files and directories across the system.

### Grant Full Disk Access to Terminal

1. Open **System Preferences** (or **System Settings** on macOS Ventura+)
2. Go to **Security & Privacy** → **Privacy** tab
3. Select **Full Disk Access** from the left sidebar
4. Click the lock icon and enter your password
5. Click the **+** button and add your Terminal application:
   - **Terminal.app**: `/Applications/Utilities/Terminal.app`
   - **iTerm2**: `/Applications/iTerm.app` (if using iTerm2)
   - **VS Code Terminal**: `/Applications/Visual Studio Code.app` (if using VS Code integrated terminal)
6. Make sure the checkbox next to your terminal app is checked
7. Restart your terminal application

### Why Full Disk Access is Required

- Leither needs to access files across the system for its container operations
- LaunchAgents have restricted access without Full Disk Access permission
- This permission allows Leither to create and manage files in various directories
- Without this permission, Leither may fail to start or function properly

## Step 2: Move Leither to an Accessible Location

macOS has privacy restrictions on the `~/Documents` folder that prevent LaunchAgents from accessing it. Move Leither to `/usr/local/darwin`:

```bash
# Copy the entire darwin directory
sudo cp -R /Users/cfa532/Documents/GitHub/darwin /usr/local/

# Set proper ownership
sudo chown -R cfa532:staff /usr/local/darwin

# Make sure the binary is executable
sudo chmod +x /usr/local/darwin/Leither

# Make directory writable (Leither needs to create PID files)
chmod -R u+w /usr/local/darwin
```

## Step 3: Test Leither from New Location

Before creating the service, verify Leither works from the new location:

```bash
# Navigate to the directory
cd /usr/local/darwin

# Run Leither manually
./Leither run

# If it starts successfully, stop it with Ctrl+C
```

## Step 4: Create the LaunchAgent Plist File

Create the plist configuration file:

```bash
nano ~/Library/LaunchAgents/com.leither.service.plist
```

Add the following content:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.leither.service</string>
    
    <key>ProgramArguments</key>
    <array>
        <string>/usr/local/darwin/Leither</string>
        <string>run</string>
    </array>
    
    <key>WorkingDirectory</key>
    <string>/usr/local/darwin</string>
    
    <key>RunAtLoad</key>
    <true/>
    
    <key>StartInterval</key>
    <integer>60</integer>
    
    <key>KeepAlive</key>
    <dict>
        <key>SuccessfulExit</key>
        <false/>
    </dict>
    
    <key>ThrottleInterval</key>
    <integer>60</integer>
    
    <key>StandardOutPath</key>
    <string>/tmp/leither.log</string>
    
    <key>StandardErrorPath</key>
    <string>/tmp/leither.err</string>
</dict>
</plist>
```

### Plist Configuration Explained

- **Label**: Unique identifier for the service
- **ProgramArguments**: Path to executable and arguments
- **WorkingDirectory**: Directory where Leither runs
- **RunAtLoad**: Start service when loaded (at login)
- **StartInterval**: Try to start every 60 seconds (prevents early boot failures)
- **KeepAlive/SuccessfulExit**: Restart if it crashes, but not if it exits cleanly
- **ThrottleInterval**: Wait 60 seconds between restart attempts
- **StandardOutPath**: Location for standard output logs
- **StandardErrorPath**: Location for error logs

## Step 5: Set Proper Permissions

```bash
chmod 644 ~/Library/LaunchAgents/com.leither.service.plist
```

## Step 6: Validate the Plist File

```bash
plutil -lint ~/Library/LaunchAgents/com.leither.service.plist
```

## Step 7: Load the Service

```bash
# Load the service
launchctl bootstrap gui/$(id -u) ~/Library/LaunchAgents/com.leither.service.plist

# Check if it's running
launchctl list | grep leither
```

Expected output:
```
3149    0    com.leither.service
```

## Step 8: Check Logs

```bash
# View standard output
cat /tmp/leither.log

# View errors
cat /tmp/leither.err

# Monitor logs in real-time
tail -f /tmp/leither.log
```

## Step 9: Test After Reboot

```bash
sudo reboot
```

After logging back in, wait about 60 seconds, then check:

```bash
launchctl list | grep leither
cat /tmp/leither.log
```

## Managing the Service

### Unload (Stop and Remove)
```bash
launchctl bootout gui/$(id -u)/com.leither.service
```

### Reload After Changes
```bash
launchctl bootout gui/$(id -u)/com.leither.service
launchctl bootstrap gui/$(id -u) ~/Library/LaunchAgents/com.leither.service.plist
```

## Troubleshooting

### Service Shows "-" Instead of PID
Check logs: `cat /tmp/leither.err`

### "Input/output error" When Loading
1. Service is already loaded (unload first)
2. Invalid plist syntax (run `plutil -lint`)

### "Operation not permitted" Errors
Executable is in protected location. Move to `/usr/local/darwin`.

### "Full Disk Access" Permission Denied
If Leither fails to start due to permission issues:

1. **Check Full Disk Access Status**:
   ```bash
   # Check if Terminal has Full Disk Access
   sqlite3 ~/Library/Application\ Support/com.apple.TCC/TCC.db "SELECT * FROM access WHERE service='kTCCServiceSystemPolicyAllFiles';"
   ```

2. **Grant Full Disk Access** (if not already done):
   - Go to **System Preferences** → **Security & Privacy** → **Privacy** → **Full Disk Access**
   - Add your Terminal application
   - Restart Terminal and try again

3. **Alternative: Use LaunchDaemon instead of LaunchAgent**:
   - Move plist to `/Library/LaunchDaemons/` (requires sudo)
   - This gives the service system-level permissions

### Service Doesn't Start After Reboot
Increase `StartInterval` to 90 or 120 seconds

## Important Notes

1. **Don't use sudo** to load LaunchAgents
2. **Use absolute paths** in the plist file
3. **The 60-second delay** prevents failures during boot

## File Locations

- **Plist**: `~/Library/LaunchAgents/com.leither.service.plist`
- **Binary**: `/usr/local/darwin/Leither`
- **Logs**: `/tmp/leither.log` and `/tmp/leither.err`

---

**Setup completed successfully! Leither will now start automatically when you log in to macOS.**
