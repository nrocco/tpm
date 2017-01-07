.DEFAULT_GOAL := test

prefix ?= /usr

VIRTUAL_ENV ?= env
PY = $(VIRTUAL_ENV)/bin/python
PIP = $(VIRTUAL_ENV)/bin/pip


$(PY):
	python3 -m venv $(VIRTUAL_ENV)


$(PIP): $(PY)
	$(PIP) install --upgrade pip
	$(PIP) install --upgrade setuptools


.PHONY: venv
venv: $(PY) $(PIP)
	$(PIP) install -r requirements.txt


.PHONY: test
test: venv
	python3 setup.py test


.PHONY: dist
dist: test
	python3 setup.py sdist


.PHONY: install
install:
	python3 setup.py install --prefix="$(prefix)" --root="$(DESTDIR)" --optimize=1


.PHONY: clean
clean:
	find * -path $(VIRTUAL_ENV) -prune -o -type d -name __pycache__ | grep __pycache__ | xargs rm -rf
	rm -rf .tox *.egg dist build .coverage MANIFEST
