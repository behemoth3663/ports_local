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

check-version: .PHONY
	find /usr/ports -mindepth 2 -maxdepth 2 -type d -name 'p5-*' | while read -r dir; do \
		local_version="$$(make -C $${dir} -V PORTVERSION)"; \
		cpan_name="$$(echo $${dir} | sed -r -e 's|.*/p5-||' -e s/-/::/g)"; \
		cpan_version="$$(fetch --quiet --output=- https://fastapi.metacpan.org/v1/module/$${cpan_name} 2>/dev/null | ${JQ_CMD} .version)"; \
		test -n "$${cpan_version}" || !	echo "$${dir##/usr/ports/} - $${cpan_name}: cannot get reply from cpan" || continue; \
		cpan_version="$${cpan_version#v}"; \
		test "$${local_version}" = "$${cpan_version}" && continue; \
		echo "$${dir##/usr/ports/} $${local_version} - $${cpan_name} $${cpan_version}"; \
	done

check-version-installed: .PHONY
	pkg query --evaluate '%n ~ p5-*' %o | while read -r origin; do \
		dir="/usr/ports/$${origin}"; \
		local_version="$$(make -C $${dir} -V PORTVERSION)"; \
		cpan_name="$$(echo $${origin} | sed -r -e 's|.*/p5-||' -e s/-/::/g)"; \
		test "$${origin}" = devel/p5-Locale-libintl && cpan_name=Locale::gettext_xs; \
		test "$${origin}" = www/p5-libwww && cpan_name=LWP; \
		cpan_version="$$(fetch --quiet --output=- https://fastapi.metacpan.org/v1/module/$${cpan_name} 2>/dev/null | ${JQ_CMD} .version)"; \
		test -n "$${cpan_version}" || !	echo "$${origin} ${local_version} - $${cpan_name}: cannot get reply from cpan" || continue; \
		cpan_version="$${cpan_version#v}"; \
		test "echo $${local_version}==$${cpan_version} | bc" = 0 || continue; \
		echo "$${origin} $${local_version} - $${cpan_name} $${cpan_version}"; \
	done
