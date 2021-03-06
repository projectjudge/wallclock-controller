#############################
# VARIABLES
#############################
START 										:= $(shell printf "\033[34;1m")
END 											:= $(shell printf "\033[0m")
SHELL 										:= /bin/bash
BUILD_SCRIPTS 						:= __build__/scripts
ENV   										?= dev
LOG_LEVEL									?= warn

#############################
# CUSTOM FUNCTIONS
#############################
define header
  $(info $(START)▶▶▶ $(1)$(END))
endef

#############################
# PHONIES
#############################
.PHONY: simple-deploy before-deploy deploy

#############################
# TARGETS
#############################
simple-deploy:
	$(call header,SIMPLE DEPLOY...)
	$(SHELL) $(BUILD_SCRIPTS)/simple_deploy.sh $(LOG_LEVEL) $(DEV_CRONJOB)

deploy: before-deploy
	$(call header,DEPLOYING...)
	$(SHELL) $(BUILD_SCRIPTS)/skaffold.sh $(LOG_LEVEL) $(DEV_CRONJOB)

before-deploy:
	$(call header,BEFORE DEPLOY...)
	$(SHELL) $(BUILD_SCRIPTS)/before_deploy.sh

# before-build:
# 	$(call header,BEFORE BUILD...)
# 	$(SHELL) $(BUILD_SCRIPTS)/before_build.sh
