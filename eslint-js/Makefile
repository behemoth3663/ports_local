PORTNAME=	js
PORTVERSION=	10.0.1
CATEGORIES=	local devel
MASTER_SITES=	LOCAL/o1vm/:npm
PKGNAMEPREFIX=	eslint-
DISTFILES=	${PORTNAME}-${DISTVERSION}${EXTRACT_SUFX}

MAINTAINER=	none@freebsd.org
COMMENT=	ESLint JavaScript Plugin
WWW=		https://eslint.org/ https://github.com/eslint/eslint

LICENSE=	MIT

FETCH_DEPENDS=	npm:www/npm \
		jq:textproc/jq \
		${LOCALBASE}/share/certs/ca-root-nss.crt:security/ca_root_nss

USES=		nodejs:run

GH_ACCOUNT=	eslint
GH_PROJECT=	js

NO_BUILD=	yes
NO_ARCH=	yes

PACKAGE_NAME=	@eslint/js

FETCH_SCRIPT=	${PORTSDIR}/Tools/scripts/npmjs-fetch-with-dependencies.sh

do-fetch: .PHONY
	@if ! [ -f ${DISTDIR}/${PORTNAME}-${DISTVERSION}${EXTRACT_SUFX} ]; then \
		${SETENV} TMPDIR=${WRKDIR} LOCALBASE=${LOCALBASE} ${FETCH_SCRIPT} \
			${PACKAGE_NAME} ${DISTVERSION} \
			${FILESDIR}/package.json \
			${DISTDIR}/${PORTNAME}-${DISTVERSION}${EXTRACT_SUFX}; \
	fi

do-install:
	@${MKDIR} ${STAGEDIR}${PREFIX}/lib
	@cd ${WRKSRC} && ${COPYTREE_SHARE} node_modules ${STAGEDIR}${PREFIX}/lib

.include <bsd.port.mk>
