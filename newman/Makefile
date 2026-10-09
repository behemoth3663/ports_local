PORTNAME=	newman
PORTVERSION=	6.2.3
CATEGORIES=	local devel
MASTER_SITES=	LOCAL/o1vm/:npm

MAINTAINER=	none@freebsd.org
COMMENT=	Command-line collection runner for Postman
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

FETCH_DEPENDS=	npm:www/npm \
		jq:textproc/jq \
		${LOCALBASE}/share/certs/ca-root-nss.crt:security/ca_root_nss

USES=		nodejs:run

GH_ACCOUNT=	postmanlabs
GH_PROJECT=	${PORTNAME}

NO_BUILD=	yes
NO_ARCH=	yes

PACKAGE_NAME=	${PORTNAME}

FETCH_SCRIPT=	${PORTSDIR}/Tools/scripts/npmjs-fetch-with-dependencies.sh

do-fetch: .PHONY
	for f in ${DISTFILES}; do \
		${TEST} -f "${DISTDIR}/$${f}" && continue; \
		${RM} -fv -- ${FILESDIR}/package.json; \
		${SETENV} TMPDIR=${WRKDIR} LOCALBASE=${LOCALBASE} \
			${FETCH_SCRIPT} ${PACKAGE_NAME} ${DISTVERSION} ${FILESDIR}/package.json "${DISTDIR}/$${f}"; \
	done

do-install: .PHONY
	@${MKDIR} ${STAGEDIR}${PREFIX}/lib
	@${FIND} ${WRKSRC}/node_modules -name '*.orig' -type f -delete
	@cd ${WRKSRC} && ${COPYTREE_SHARE} node_modules ${STAGEDIR}${PREFIX}/lib
	@${CHMOD} +x ${STAGEDIR}${PREFIX}/lib/node_modules/${PACKAGE_NAME}/bin/${PACKAGE_NAME}.js
	@${MKDIR} ${STAGEDIR}${PREFIX}/bin
	@${RLN} ${STAGEDIR}${PREFIX}/lib/node_modules/${PACKAGE_NAME}/bin/${PACKAGE_NAME}.js ${STAGEDIR}${PREFIX}/bin/${PACKAGE_NAME}

.include <bsd.port.mk>
