.PHONY: secret

create-secret:
	kubectl crreate secret generic ${n} ${o}
