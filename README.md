# manitool
Tool for Kubernetes manifests

For example, you might use the sed command when you want to rewrite one line of a file inside the initContainers in the Kubernetes manifest file. Have you ever had any trouble writing in a YAML file because the string to be replaced contains single or double quotes?
This tool solves these problems.

Since this tool is written in golang, it generates very small container images.
You can get the generated container as follows:
docker pull tlj9if8lmhrr / manitool

[How to use]

manitool replace [file path] [single quote symbol] [double quote symbol] [old string] [new string]

If you have a text file named /etc/sample.conf contains a line like:

#parameter1 'x, y'

then you can replece this line to

parameter1 'x, "z"'

by using this tool as

manitool replace /etc/sample.conf @@ %% '#parameter1 @@x, y@@' 'parameter1 @@x, %%z%%@@'
