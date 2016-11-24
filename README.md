# Hushtache

![Hushtache Logo](//static.hushtache.com/logo.png)

Config management with built-in encryption that allow developers to store app secrets encrypted in the same repo. Secrets are encrypted using AES that has the key protected using a generated private key of each user.

Along with simple storing and retrieving secret, the tool allows rendering directing to any templates given. Examples of these might include Google App Engine's app.yaml, where the environment variables are populated from the secret store.

## Install

To install download the latest build for your platform:

* [MacOS 64](https://builds.hushtache.com/20161124/darwin64.zip)
* [MacOS 32](https://builds.hushtache.com/20161124/darwin32.zip)
* [Linux 64](https://builds.hushtache.com/20161124/linux64.zip)
* [Linux 32](https://builds.hushtache.com/20161124/linux32.zip)
* [Windows 64](https://builds.hushtache.com/20161124/windows64.zip)
* [Windows 32](https://builds.hushtache.com/20161124/windows32.zip)

THe zip will contain the executable that can be used directly or can be added to your global path using:

```
mv hushtache /usr/bin/hushtache
```

## Start

To start hushtache needs to generate a secret store where the secrets can added to, much like `git` the command to init in a folder is:

```
hushtache init
```

To check if you have have access and able to decrypt the store in the current directory, use:

```
hushtache allowed
```

## Secrets

Secrets can then be set (or updated) using:

```
hushtache set <key> <value>
```

And all displayed using:
```
# show all keys
hushtache get

# show key value
hushtache get <key>
```

Secrets can also be removed using:

```
hushtache remove <key>
```

## Rendering

The idea with the local store is to make configuration files easy to generate without storing the actual secrets inside of the repo itself in plaintext.

The templating language `Handlebars` is built in and can be called using:

```
hushtache render <location-to-file> <location-to-file2>
```

Which will output the generated file to STDOUT ready for use. An example of using this could be Google App Engine that makes use of a app.yaml. 

A file named `app.template` could be in the root of the repo:

```
runtime: python
environment_variables:
  PORT: {{PORT}}
  SECRET_TOKEN: {{SECRET_TOKEN}}
```

The secrets can then be set in the store:

```
hushtache set PORT 8080
hushtache set SECRET_TOKEN supersecret
```

And render called to populate the template from the encrypted store:

```
hushtache render app.template
```

Which would produce:

```
runtime: python
environment_variables:
  PORT: 8080
  SECRET_TOKEN: supersecret
```

> Undefined keys will be populated with a empty string

## Users

Once created, users who are allowed to access the secrets can be seen using:

```
hushtache users
```

More users can be added by asking the users to provide the output of `hushtache key`, which is their public key. This user can then be added to allow decryption and access to the store using:

```
hushtache allow <username> <publickey>
```

Users can also be removed:
```
hushtache deny <username>
```

# License 

Copyright 2016 Hushtache

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.