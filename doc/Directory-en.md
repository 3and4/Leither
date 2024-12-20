Root Directory <div id="rootdir"></div> 
｜Leither-----------------System Program   
｜SystemVars.json---------[Configuration File](#systemvars)  
｜hostkey.cfg-------------[Node Key](#hostkey)  
｜temp.html---------------[Domain Template](#temp)  
｜Leither.log-------------Log File  
｜--filestore-------------Mapped to Hard Drive or External Storage  
｜--mac-------------------Stores Basic Files, Named by Hash  
｜--mid-------------------Stores Media Files, Named by Media ID  
｜--service---------------Service Directory, Stores Server Scripts  
｜   ｜----RequestService--Script for Handling Service Requests  
｜--webdav----------------Default WebDAV Directory     
｜--db--------------------Database Directory  
｜   ｜----backup----------Static Backup of Media Database  
｜   ｜----mmdb------------Dynamic Data of Media Database  
｜--temp------------------Temporary Files Directory  
  
  
### Configuration File SystemVars.json<div id="systemvars"></div>
The file format is JSON.  
**ServicePort**  
Service Port  
If not set, the system defaults to 9000. For home nodes, ports 80 and 8080 cannot be set as they are usually blocked by ISPs.  
Configuration Example:  
"ServicePort":80,

**FixedAddr**   
Domain or Fixed IP  
Used for cloud hosts with a dedicated IP.  
If not set, the program will automatically detect its external address through the gateway or peer nodes.  
Configuration Example:  
"FixedAddr":"121.44.244.138",

**LogConfig**  
Log Configuration  
Configuration Example:  
1. File Log  
"LogConfig":"{\"filename\":\"Leither.log\",\"level\":1}",  
2. Console Log  
"console":"{\"level\":1}",  

**LAN**  
In the case of multiple network cards, binds the service to a specific network card.  
Configuration Example:   
1. Bind by MAC  
"LAN":"98-5F-D3-46-17-A8",

2. Bind by IP  
"LAN":"192.168.1.24",

**Gateway**  
Gateway Nodes  
"18.222.243.60"  
"121.40.244.135"   


### Node Key hostkey.cfg<div id="hostkey"></div>  
Stores the key for the current node.  
If this file does not exist on the first run, the system will automatically generate one.  
When executing the Leither lssl command in the current directory without specifying a key file, this file is used by default, indicating that the command is executed as the node system user.  

### Domain Template temp.html<div id="temp"></div>  
Used for domain resolution, only effective for domain nodes.  
When assisting other users in domain resolution, the node reads the relevant information corresponding to the domain, writes it into the template, generates an HTML file, and returns it. When the browser parses this HTML file, it will guide to the user-specified node to execute the relevant applications or data on the node.
