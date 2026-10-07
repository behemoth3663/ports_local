PORTNAME=	cloudburn
PORTVERSION=	0.18.9
CATEGORIES=	local sysutils
MASTER_SITES=	LOCAL/o1vm/:npm

MAINTAINER=	none@freebsd.org
COMMENT=	Policy engine that blocks bad AWS spending patterns
WWW=		https://cloudburn.io/ https://github.com/towardsthecloud/cloudburn

LICENSE=	APACHE20

FETCH_DEPENDS=	npm:www/npm \
		jq:textproc/jq \
		${LOCALBASE}/share/certs/ca-root-nss.crt:security/ca_root_nss

USES=		nodejs:run

GH_ACCOUNT=	towardsthecloud

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
	@cd ${WRKSRC} && ${COPYTREE_SHARE} node_modules ${STAGEDIR}${PREFIX}/lib
	@${CHMOD} +x ${STAGEDIR}${PREFIX}/lib/node_modules/${PACKAGE_NAME}/dist/cli.js
	@${RLN} ${STAGEDIR}${PREFIX}/lib/node_modules/${PACKAGE_NAME}/dist/cli.js ${STAGEDIR}${PREFIX}/bin/${PACKAGE_NAME}

.include <bsd.port.mk>
