## Gowie Architect Test Drive -- Notes / Feedback

Note: Please excuse my stream of consciousness style notes / feedback. This was done during a day of mostly meetings and I didn't get to devote as much time as I wanted to.

1. https://cln.sh/1JuN0C -- You should add to this CLI example that you can deploy to different 'types' (GCP, Azure) since you call that out in the description tect.
1. https://cln.sh/WESDjF -- "OIDC compliant use authentication service" => "OIDC compliant user authentication service"
1. I keep seeing different `architect.json` file names in examples: arc-environment.json, architect.json, arc-env.json, env-config.json. I'd consolidate all examples to use one or two so as to not confuse their purpose... though maybe I'm misunderstanding and these aren't all the same?
1. Personal Preference: YML > JSON. Surfacing this again because I feel like the wider community shares my opinion (k8s, helm, sops, CloudFormation, cloud-init, etc etc). If I was evaluating this tool for use at my company, there would be a small part of me at the back of my head which was asking "Why am I editing json"? Okay, I'll shut up now.
   1. Aha... it seems you already made this swtich recently but haven't updated things everywhere? Hurray!
1. https://cln.sh/ZhfkdO -- What IAM perms does this Access Key need? I gave it Administrator cause this is my biz test account, but I'd like know if I can lock this down to certain permissions for POLP.
1. Side note -- Architect is a great name for this product. Well done.
1. I hit the back button while going through the demo flow (specifically trying to go back to the "Select Hello World vs Google Microservices" demo page) and it seems to have logged me out and kicked me all the way back to the login screen. Worth investigating that one.
1. Terraform Error: https://app.architect.io/gowiem/environments/fusionauth-example/ -- "Error: InvalidParameterException: The target group with targetGroupArn arn:aws:elasticloadbalancing:us-west-2:115843287071:targetgroup/arc-20200604191342029100000005/e2a494f3bba19a53 does not have an associated load balancer. "react-demo-frontend-latest_svc""
   1. I've run into this a bunch of times on client projects as it's a difficult dependency for Terraform / the AWS provider to resolve. You need to add the ALB as an explicit dependency to your aws_alb_target_group via `depends_on`. `depends_on = [ aws_lb._architect_ecs_lb ]` might do it. You might also need to involve your ECS Service deploy as an explicit dependency as well, but not sure without looking more at the setup / tf files.
      1. Old TF issue on this: https://github.com/hashicorp/terraform/issues/12634
   1. More general feedback on this: How am I supposed to recover as an engineer / DevOps owner? Ping the Architect team? Is there anyway to get to my project's underlying TF files / state?
      1. Reapplying the same infrastructure did a `terraform -auto-apply -destroy` (or seems to have) when that was not my intention... This seems scary and I'd be put off on not having insight / approval on that terraform plan if this was a production app.
1. Usability item: I'd update the apply console logging to not auto-scroll on new log entries. This is painful when trying to read further up in the console logs.
1. Saw you're using NAT Gateways for Dev Environments. Would be great to be able to configure these types of things per environment. I typically run NAT Instances (t3.micros) in Dev / Stage environments if my clients are cost conscious, so I don't end up costing them the $32 / month per gateway + the *insane* per gb costs.
1. Similarly, what if my organization has resource tagging policies that I need my application to conform to? This is common in enterprise / larger orgs, so I imagine this would quickly be somethign you'd run up against with those types of customers.
1. Have you thought about generating the tf files for the customer and then letting them handle deployments? What can I do if I want to fold in other AWS components (WAF comes to mind) into my infra via terraform?
1. Usability item: What does "Test Connection"? This tells me I have a successful connections, but I don't actually have the app deployed so I'm not sure what is connected successfully.
1. Thought: Personally, I've seen Golang emerging as one of the major languages where its developers are shipping microservices. I'd create some examples in Golang to help showcase usage in that realm.
1. I think there could be a lot more to be said about the private vs public registry. I didn't really understand that from reading the website, but now that I'm reading the docs more I'm seeing that is a thing that should be understood. Things like `architect-community/postgres: 11` and others tipped me off, but I didn't see anything explicit surrounding this.
1. "An optional reference to the language used for the service. Including this allows Architect to generate client code for downstream dependencies and ingest them via corresponding SDKs for easier connectivity to dependencies." -- Is this still a thing? I'm seeing it in the service config ref, but didn't see it elsewhere.
1. Two service configs in one directory... Is this possible? Different file name? I'm going to give that a shot.
1. Reading architect.yml parameters in from .env or local configuration? Maybe this is against the point?
1. How can I find out what values are available for the "derived" parameters? I see URL is available and I use that... Is port available to since it's specified in the service config?
1. Really dig the `architect services` table output format. You should also list tags when specifying a particular service as that was what I originally was looking for when usign that command.
1. `architect whoami` results in: `›   Error: command whoami not found`
1. You folks should do a video walkthrough or blog post walk through of architect'ing a simple Docker application. It would be useful to see the various architect commands and their ordering.
1. Ya'll should create the $ENV.env.yml file for me when invoking `env:create`.
1. The *-sdk repos are all gone -- SDKs as an idea are gone I'm guessing?
1. `architect register --help` => "  -s, --services=services        Path to a service to build" -- Am I passing multiple service paths or just one? The help text is confusing.
1. Error: https://cln.sh/hgL7t1 -- Not sure what is up with that...
   1. Am I using an outdated version of the CLI tool or is this legacy? Same thing happened again with another service config: https://cln.sh/Q7ByDS
1. Pushing a custom service with a Dockerfile to the registry stalled: https://cln.sh/rxiIMg + https://cln.sh/Q7ByDS
   1. It was like this for a few minutes before I canceled the process.
1. Ya'll should alias `architect` to `arc`. Who wants to type that many letters?

## Wrapping Here
I stopped at this point cause I couldn't push my image and I've hit a bit of a timebox. Cool stuff so far, but obviously a lot of docs are in flux and that hindered me a bit. I think you need a walkthrough, but I'm sure you're holding off on executing that until things have solidified more which I totaly get.

## Insider questions cause I'm curious
1. Where are your worker nodes running? Your own k8s cluster? What Cloud?
1. What are you using for the backing container registry? ECR? Artifactory? Tell me not quay.