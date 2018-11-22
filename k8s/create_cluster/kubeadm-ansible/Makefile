SSH_KEY := $(HOME)/.ssh/gcloud
RUN_OPTS := --private-key=$(SSH_KEY)
RUN := ansible-playbook $(RUN_OPTS)

.PHONY: ubuntu-install centos-install master workers clean

ubuntu-install:
	$(RUN) kube-dependencies-ubuntu.yaml

centos-install:
	$(RUN) kube-dependencies-centos.yaml

master:
	$(RUN) master-node.yaml

workers:
	$(RUN) worker-nodes.yaml

clean:
	rm -rf kube-dependencies-ubuntu.retry
	rm -rf kube-dependencies-centos.retry
	rm -rf master-node.retry
	rm -rf worker-nodes.retry
