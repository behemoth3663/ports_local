PORTNAME=	eslint
PORTVERSION=	10.4.1
#PORTREVISION=	1
CATEGORIES=	local devel
MASTER_SITES=	LOCAL/o1vm/:npm
DISTFILES=	${PORTNAME}-${DISTVERSION}${EXTRACT_SUFX}

MAINTAINER=	none@freebsd.org
COMMENT=	Find and fix problems in JavaScript code
WWW=		https://eslint.org/

LICENSE=	MIT

FETCH_DEPENDS=	npm:www/npm \
		jq:textproc/jq \
		${LOCALBASE}/share/certs/ca-root-nss.crt:security/ca_root_nss

USES=		nodejs:run

GH_ACCOUNT=	eslint

NO_BUILD=	yes
NO_ARCH=	yes

PACKAGE_NAME=	${PORTNAME}

FETCH_SCRIPT=	${PORTSDIR}/Tools/scripts/npmjs-fetch-with-dependencies.sh

do-fetch: .PHONY
.for _f in ${DISTFILES}
	@test -f ${DISTDIR}/${_f} || \
		${SETENV} TMPDIR=${WRKDIR} LOCALBASE=${LOCALBASE} ${FETCH_SCRIPT} \
			${PACKAGE_NAME} ${DISTVERSION} \
			${FILESDIR}/package.json \
			${DISTDIR}/${_f}
.endfor

do-install: .PHONY
	@${MKDIR} ${STAGEDIR}${PREFIX}/lib
	@cd ${WRKSRC} && ${COPYTREE_SHARE} node_modules ${STAGEDIR}${PREFIX}/lib
	@${CHMOD} +x ${STAGEDIR}${PREFIX}/lib/node_modules/${PACKAGE_NAME}/bin/eslint.js
	@${MKDIR} ${STAGEDIR}${PREFIX}/bin
	@${RLN} ${STAGEDIR}${PREFIX}/lib/node_modules/${PACKAGE_NAME}/bin/eslint.js ${STAGEDIR}${PREFIX}/bin/${PORTNAME}
	${REINPLACE_CMD} -i '' \
		-e "s|#!/usr/bin/env[[:blank:]].*|#!${SETENV} -S NODE_PATH=${PREFIX}/lib/node_modules node|" \
		${STAGEDIR}${PREFIX}/lib/node_modules/${PACKAGE_NAME}/bin/eslint.js

.include <bsd.port.mk>
