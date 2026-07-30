PORTNAME=	npm-run-all
PORTVERSION=	4.1.5
PORTREVISION=	1
CATEGORIES=	local devel
MASTER_SITES=	LOCAL/o1vm/:npm

MAINTAINER=	none@freebsd.org
COMMENT=	Run multiple npm-scripts in parallel or sequential
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT

FETCH_DEPENDS=	npm:www/npm \
		jq:textproc/jq \
		${LOCALBASE}/share/certs/ca-root-nss.crt:security/ca_root_nss

USES=		nodejs:run

GH_ACCOUNT=	mysticatea
GH_PROJECT=	${PORTNAME}

NO_BUILD=	yes
NO_ARCH=	yes

PACKAGE_NAME=	${PORTNAME}

FETCH_SCRIPT=	${PORTSDIR}/Tools/scripts/npmjs-fetch-with-dependencies.sh

do-fetch: .PHONY .SILENT
.for _f in ${DISTFILES}
	test -f ${DISTDIR}/${_f} || ${SETENV} TMPDIR=${WRKDIR} LOCALBASE=${LOCALBASE} \
		${FETCH_SCRIPT} ${PACKAGE_NAME} ${DISTVERSION} ${FILESDIR}/package.json ${DISTDIR}/${_f}
.endfor

do-install: .PHONY
	@${MKDIR} ${STAGEDIR}${PREFIX}/lib
	@cd ${WRKSRC} && ${COPYTREE_SHARE} node_modules ${STAGEDIR}${PREFIX}/lib
	@${MKDIR} ${STAGEDIR}${PREFIX}/bin
	@${RLN} ${STAGEDIR}${PREFIX}/lib/node_modules/npm-run-all/bin/npm-run-all/index.js ${STAGEDIR}${PREFIX}/bin/${PACKAGE_NAME}
	@${RLN} ${STAGEDIR}${PREFIX}/lib/node_modules/npm-run-all/bin/run-p/index.js ${STAGEDIR}${PREFIX}/bin/run-p
	@${RLN} ${STAGEDIR}${PREFIX}/lib/node_modules/npm-run-all/bin/run-s/index.js ${STAGEDIR}${PREFIX}/bin/run-s
	@${CHMOD} +x ${STAGEDIR}${PREFIX}/lib/node_modules/npm-run-all/bin/npm-run-all/index.js
	@${CHMOD} +x ${STAGEDIR}${PREFIX}/lib/node_modules/npm-run-all/bin/run-p/index.js
	@${CHMOD} +x ${STAGEDIR}${PREFIX}/lib/node_modules/npm-run-all/bin/run-s/index.js

.include <bsd.port.mk>
