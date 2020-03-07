# manitool
Tool for Kubernetes manifests

For example, you might use the sed command when you want to rewrite one line of a file inside the initContainers in the Kubernetes manifest file. Have you ever had any trouble writing in a YAML file because the string to be replaced contains single or double quotes?  
This tool solves these problems.

Since this tool is written in golang, it generates very small container images.  
You can get the generated container as follows
```
docker pull tlj9if8lmhrr/manitool
```

## How to use
```
manitool replace [file path] [single quote symbol] [double quote symbol] [old string] [new string]
```

If you have a text file named /etc/sample.conf contains a line like
```
#parameter1 'x, y'
```
then you can replece this line to
```
parameter1 'x, "z"'
```
by using this tool as  
```
manitool replace /etc/sample.conf @@ %% '#parameter1 @@x, y@@' 'parameter1 @@x, %%z%%@@'
```

### HOSTNAME environment variable support
You can include the value of $HOSTNAME by specifying two addtional options.
```
manitool replace [file path] [single quote symbol] [double quote symbol] [old string] [new string] [hostname symbol] [hyphen alt ]
```

Assuming the $HOSTNAME is 'name-0' and if you execute the following command
```
manitool replace /etc/sample.conf @@ %% '#parameter1 @@x, y@@' 'parameter1 @@x, %%z%%, {{HOSTNAME}}@@' {{HOSTNAME}} _
```
then you will get the following string where the hostname 'name-0' is replaced to 'name_0' 
```
parameter1 'x, "z", name_0'
```
