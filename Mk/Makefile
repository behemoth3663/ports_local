JQ_CMD?=	/usr/local/bin/jq --raw-output --monochrome-output

.MAIN: all

all: .PHONY builtin.json site_perl.json

clean: .PHONY
	rm -f -- builtin.json site_perl.json

#	( make MAKE_JSON_DIRS="$$(perl -e 'print((sort @INC)[0])')" make-json  )
make-json: .PHONY .SILENT
	test -n '${MAKE_JSON_DIRS}' && find ${MAKE_JSON_DIRS} -type f -name '*.pm' -exec \
		grep --extended-regexp '^package[[:space:]]+[^;]+[[:space:]]*;[[:space:]]*$$' {} + | \
		sed -rn 's/^(.*):package[[:space:]]+([^;]+).*$$/"\2": "\1",/p' | \
		sort | sed -e '1 s/^/{/' -e '$$ s/,$$/}/' | \
		jq --raw-output --monochrome-output .
	test -n '${MAKE_JSON_DIRS}' && find ${MAKE_JSON_DIRS} -type f -name '*.pm' -exec \
		sed -rn 's/^package[[:space:]]+([^[:space:];#]+)[^#]*;.*$$/\1/p' {} + | sort --unique | while read -r pkg; do \
		echo "use $${pkg}; print \$$$${pkg}::VERSION"; \
		test $${pkg} = _charnames -o $${pkg} = '$$pack' -o $${pkg} = '$$this_package' || ver=$$(perl -e "use '$${pkg}'; print \$$$${pkg}::VERSION" || :); \
		test -z "$${ver}" || echo "$${pkg} $${}" >&2; \
	done

builtin.json: .PHONY
	perl -e ' \
		use strict; \
		use warnings; \
		use integer; \
		use Module::CoreList; \
		use JSON; \
		print JSON->new->canonical->encode($$Module::CoreList::version{$$]}); \
	' | ${JQ_CMD} . > ${.TARGET}


site_perl.json: .PHONY
	perl -e ' \
		use strict; \
		use warnings; \
		use integer; \
		use ExtUtils::Installed; \
		use JSON; \
		my $$installed = ExtUtils::Installed->new(); \
		my %modules = map { $$_ => $$installed->version($$_) } $$installed->modules(); \
		print JSON->new->canonical->encode(\%modules); \
	' | ${JQ_CMD} . > ${.TARGET}
