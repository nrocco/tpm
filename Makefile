.DEFAULT_GOAL := test

prefix ?= /usr

VIRTUAL_ENV ?= env
PY = $(VIRTUAL_ENV)/bin/python
PIP = $(VIRTUAL_ENV)/bin/pip
PASSCTL = $(VIRTUAL_ENV)/bin/passctl
PYTEST = $(VIRTUAL_ENV)/bin/pytest
TABULATE = $(VIRTUAL_ENV)/bin/tabulate


$(PY):
	python3 -m venv $(VIRTUAL_ENV)


$(TABULATE): $(PY)
	$(PIP) install -r requirements.txt


$(PYTEST): $(PY)
	$(PIP) install -r requirements-dev.txt


$(PASSCTL): $(PY) $(TABULATE)
	$(PY) setup.py develop


.PHONY: venv
venv: $(PY) $(TABULATE)


.PHONY: develop
develop: $(PASSCTL)


.PHONY: test
test: $(PASSCTL) $(PYTEST)
	$(PYTEST) --cov=team_password_cli --no-cov-on-fail --cov-report term --cov-report html tests/ $(ARGS)


.PHONY: dist
dist:
	$(PY) setup.py sdist


.PHONY: install
install:
	$(PY) setup.py install --prefix="$(prefix)" --root="$(DESTDIR)" --optimize=1


.PHONY: clean
clean:
	find * -path $(VIRTUAL_ENV) -prune -o -type d -name __pycache__ | grep __pycache__ | xargs rm -rf
	rm -rf .tox *.egg dist build .coverage MANIFEST
