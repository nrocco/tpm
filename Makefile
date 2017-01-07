.DEFAULT_GOAL := test

prefix ?= /usr

VIRTUAL_ENV ?= env
PY = $(VIRTUAL_ENV)/bin/python
PIP = $(VIRTUAL_ENV)/bin/pip


$(PY):
	python3 -m venv $(VIRTUAL_ENV)


.PHONY: venv
venv: $(PY)
	$(PIP) install -r requirements.txt


.PHONY: test
test: venv
	$(PY) setup.py test


.PHONY: dist
dist: test
	$(PY) setup.py sdist


.PHONY: install
install: venv
	$(PY) setup.py install --prefix="$(prefix)" --root="$(DESTDIR)" --optimize=1


.PHONY: clean
clean:
	find * -path $(VIRTUAL_ENV) -prune -o -type d -name __pycache__ | grep __pycache__ | xargs rm -rf
	rm -rf .tox *.egg dist build .coverage MANIFEST
