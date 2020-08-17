#https://www.freebsd.org/doc/en_US.ISO8859-1/books/porters-handbook/building.html#using-go
#rm -f -- distinfo Makefile.deps && make makesum && make go113mod-vendor && make makesum && mkcal
#GOFLAGS:=${GOFLAGS:S/-modcacherw/}

${WRKSRC}/.go113mod-vendor:
	${MKDIR} ${WRKDIR}
	cd ${WRKDIR} && for f in ${EXTRACT_ONLY}; do \
		${EXTRACT_CMD} ${EXTRACT_BEFORE_ARGS} ${_DISTDIR}/$${f} ${EXTRACT_AFTER_ARGS}; \
	done
	touch ${WRKSRC}/.go113mod-vendor

${WRKSRC}/vendor/modules.txt: ${WRKSRC}/.go113mod-vendor
	cd ${WRKSRC}; \
	${TEST} -f go.mod || ${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod init ${GO_PKGNAME}; \
	${SETENV} GOPATH=${WRKDIR}/.gopath ${GO_CMD} mod vendor

${.CURDIR}/Makefile.deps: ${WRKSRC}/vendor/modules.txt
	modules2tuple ${WRKSRC}/vendor/modules.txt >${.CURDIR}/Makefile.deps

go113mod-vendor: .PHONY ${.CURDIR}/Makefile.deps
	${CHMOD} -R u+w ${WRKDIR}
	${RM} -r -- ${WRKDIR} ${DISTINFO_FILE}
