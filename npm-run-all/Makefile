PORTNAME=	npm-run-all
PORTVERSION=	4.1.5
CATEGORIES=	local devel
MASTER_SITES=	LOCAL/o1vm/:npm
DISTFILES=	${PORTNAME}-${DISTVERSION}${EXTRACT_SUFX}

MAINTAINER=	none@freebsd.org
COMMENT=	Run multiple npm-scripts in parallel or sequential
WWW=		https://github.com/mysticatea/npm-run-all

LICENSE=	MIT

FETCH_DEPENDS=	npm:www/npm \
		jq:textproc/jq \
		${LOCALBASE}/share/certs/ca-root-nss.crt:security/ca_root_nss

USES=		nodejs:run

GH_ACCOUNT=	mysticatea
GH_PROJECT=	npm-run-all

NO_BUILD=	yes
NO_ARCH=	yes

PACKAGE_NAME=	${PORTNAME}

FETCH_SCRIPT=	${PORTSDIR}/Tools/scripts/npmjs-fetch-with-dependencies.sh

do-fetch: .PHONY
	@if ! [ -f ${DISTDIR}/${PORTNAME}-${DISTVERSION}${EXTRACT_SUFX} ]; then \
		${SETENV} TMPDIR=${WRKDIR} LOCALBASE=${LOCALBASE} ${FETCH_SCRIPT} \
			${PACKAGE_NAME} ${DISTVERSION} \
			${FILESDIR}/package.json \
			${DISTDIR}/${PORTNAME}-${DISTVERSION}${EXTRACT_SUFX}; \
	fi

do-install: .PHONY
	@${MKDIR} ${STAGEDIR}${PREFIX}/lib
	@cd ${WRKSRC} && ${COPYTREE_SHARE} node_modules ${STAGEDIR}${PREFIX}/lib
	@${MKDIR} ${STAGEDIR}${PREFIX}/bin
	@${RLN} ${STAGEDIR}${PREFIX}/lib/node_modules/npm-run-all/bin/npm-run-all/index.js ${STAGEDIR}${PREFIX}/bin/npm-run-all
	@${RLN} ${STAGEDIR}${PREFIX}/lib/node_modules/npm-run-all/bin/run-p/index.js ${STAGEDIR}${PREFIX}/bin/run-p
	@${RLN} ${STAGEDIR}${PREFIX}/lib/node_modules/npm-run-all/bin/run-s/index.js ${STAGEDIR}${PREFIX}/bin/run-s
	@${CHMOD} +x ${STAGEDIR}${PREFIX}/lib/node_modules/npm-run-all/bin/npm-run-all/index.js
	@${CHMOD} +x ${STAGEDIR}${PREFIX}/lib/node_modules/npm-run-all/bin/run-p/index.js
	@${CHMOD} +x ${STAGEDIR}${PREFIX}/lib/node_modules/npm-run-all/bin/run-s/index.js

.include <bsd.port.mk>
