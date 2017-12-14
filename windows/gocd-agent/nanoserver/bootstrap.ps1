$properties = @("agent.auto.register.key=${$env:GO_EA_AUTO_REGISTER_KEY}",
"agent.auto.register.environments=${$env:GO_EA_AUTO_REGISTER_ENVIRONMENT}",
"agent.auto.register.elasticAgent.agentId=${$env:GO_EA_AUTO_REGISTER_ELASTIC_AGENT_ID}",
"agent.auto.register.elasticAgent.pluginId=${$env:GO_EA_AUTO_REGISTER_ELASTIC_PLUGIN_ID}")

@("config", "work", "logs") | % { New-Item -ItemType Directory C:\gocd-agent\$_ }

$properties | Out-File "C:\gocd-agent\config\autoregister.properties"

$env:GO_SERVER_URL = $env:GO_EA_SERVER_URL
$env:GOCD_AGENT_WORKING_DIR = "C:\gocd-agent\work"
$env:GOCD_AGENT_LOG_DIR = "C:\gocd-agent\logs"

& C:\gocd-agent\agent.exe