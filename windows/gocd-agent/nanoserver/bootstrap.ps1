New-Item -ItemType Directory .\config

$properties = @("agent.auto.register.key=$env:GO_EA_AUTO_REGISTER_KEY",
"agent.auto.register.environments=$env:GO_EA_AUTO_REGISTER_ENVIRONMENT",
"agent.auto.register.elasticAgent.agentId=$env:GO_EA_AUTO_REGISTER_ELASTIC_AGENT_ID",
"agent.auto.register.elasticAgent.pluginId=$env:GO_EA_AUTO_REGISTER_ELASTIC_PLUGIN_ID")

$properties | Out-File "C:\gocd-agent\config\autoregister.properties" -Encoding "default" -append

$env:GO_SERVER_URL=$env:GO_EA_SERVER_URL

& .\agent.cmd