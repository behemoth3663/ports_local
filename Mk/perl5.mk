JQ_CMD?=	/usr/local/bin/jq --raw-output --monochrome-output

get-depends: .PHONY .SILENT
		while read -r s; do \
			set -- $${s}; \
			test "$${1}" = perl && continue; \
			test "$$(${JQ_CMD} .\"$${1}\" /usr/ports/local/Mk/builtin.json)" = null || continue; \
			port_name="p5-$$(echo $${1} | sed -r s/:+/-/g)"; \
			port_dir=$$(find /usr/ports -mindepth 2 -maxdepth 2 -type d -name "$${port_name}" | sed s,^/usr/ports/,,); \
			if [ -n "$${port_dir}" ]; then \
				echo "	$${port_name}>=$${2}:$${port_dir}"; \
			else \
				fetch --quiet --output=- "https://fastapi.metacpan.org/v1/module/$${1}" | ${JQ_CMD} .download_url; \
			fi; \
		done

get-build-depends: .PHONY patch
		${JQ_CMD} '.prereqs.build | select(.requires != null) | .requires | to_entries | .[] | .key + " " + .value' ${WRKSRC}/META.json | \
		( make -C ${.CURDIR} get-depends )

get-run-depends: .PHONY .SILENT patch
		${JQ_CMD} '.prereqs.runtime | select(.requires != null) | .requires | to_entries | .[] | .key + " " + .value' ${WRKSRC}/META.json | \
		( make -C ${.CURDIR} get-depends )

get-test-depends: .PHONY .SILENT patch
		${JQ_CMD} '.prereqs.test | select(.requires != null) | .requires | to_entries | .[] | .key + " " + .value' ${WRKSRC}/META.json | \
		( make -C ${.CURDIR} get-depends )
