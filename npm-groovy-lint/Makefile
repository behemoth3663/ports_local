PORTNAME=	npm-groovy-lint
PORTVERSION=	19.0.1
CATEGORIES=	local devel
MASTER_SITES=	LOCAL/o1vm/:npm

MAINTAINER=	none@freebsd.org
COMMENT=	Lint, format and auto-fix Groovy, Jenkinsfile, Gradle files
WWW=		https://nvuillam.github.io/npm-groovy-lint/ https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

FETCH_DEPENDS=	npm:www/npm \
		jq:textproc/jq \
		${LOCALBASE}/share/certs/ca-root-nss.crt:security/ca_root_nss

USES=		nodejs:run

GH_ACCOUNT=	nvuillam
GH_PROJECT=	${PORTNAME}

NO_BUILD=	yes
NO_ARCH=	yes

PACKAGE_NAME=	${PORTNAME}

FETCH_SCRIPT=	${PORTSDIR}/Tools/scripts/npmjs-fetch-with-dependencies.sh

do-fetch: .PHONY
	DISTFILE=${PORTNAME}-${PORTVERSION}${EXTRACT_SUFX}; \
	if [ ! -f ${DISTDIR}/$${DISTFILE} ]; then \
		${RM} -fv -- ${FILESDIR}/package.json; \
		${SETENV} TMPDIR=${WRKDIR} LOCALBASE=${LOCALBASE} \
			${FETCH_SCRIPT} ${PACKAGE_NAME} ${PORTVERSION} ${FILESDIR}/package.json ${DISTDIR}/$${DISTFILE}; \
	fi

do-install: .PHONY
	@${MKDIR} ${STAGEDIR}${PREFIX}/lib
	@cd ${WRKSRC} && ${COPYTREE_SHARE} node_modules ${STAGEDIR}${PREFIX}/lib
	@${CHMOD} +x ${STAGEDIR}${PREFIX}/lib/node_modules/${PACKAGE_NAME}/lib/index.js
	@${MKDIR} ${STAGEDIR}${PREFIX}/bin
	@${RLN} ${STAGEDIR}${PREFIX}/lib/node_modules/${PACKAGE_NAME}/lib/index.js ${STAGEDIR}${PREFIX}/bin/${PACKAGE_NAME}

.include <bsd.port.mk>
