PORTNAME=	js
PORTVERSION=	10.0.1
PORTREVISION=	1
CATEGORIES=	local devel
MASTER_SITES=	LOCAL/o1vm/:npm
PKGNAMEPREFIX=	eslint-
DISTFILES=	${PORTNAME}-${PORTVERSION}${EXTRACT_SUFX}

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
	test -f ${DISTDIR}/${DISTFILES} || ${SETENV} TMPDIR=${WRKDIR} LOCALBASE=${LOCALBASE} \
		${FETCH_SCRIPT} ${PACKAGE_NAME} ${PORTVERSION} ${FILESDIR}/package.json ${DISTDIR}/${DISTFILES}

do-install: .PHONY
	${MKDIR} ${STAGEDIR}${PREFIX}/lib
	cd ${WRKSRC} && ${COPYTREE_SHARE} node_modules ${STAGEDIR}${PREFIX}/lib

.include <bsd.port.mk>
