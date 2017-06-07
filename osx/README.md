# Cider

A command line utility to manage Gauge GoCD build [OSX agent](https://atlas.hashicorp.com/getgauge).

## Usage

Set the following environment variables

* `AGENT_AUTO_REGISTER_KEY`
* `AGENT_NAME`
* `GAUGE_DOWNLOADS_IP`
* `GO_SERVER_URL`

Run the following command to start the agent.
```
cider <box name> <IP>
```

The above command will perform the following actions:

* Install vagrant `sahara` plugin for sandboxing.
* Start the vagrant VM.
* Set Go server auto register properties.
* Start Go Agent.
* Watch for rollbacks. It watches for `rollback.txt` in `/vagrant` folder. If that file is found, it will clean up the machine and bring back to its original state.
