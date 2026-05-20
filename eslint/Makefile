PORTNAME=	eslint
PORTVERSION=	10.4.0
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

NO_BUILD=	yes
NO_ARCH=	yes

PACKAGE_NAME=	eslint

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
	@${CHMOD} +x ${STAGEDIR}${PREFIX}/lib/node_modules/${PACKAGE_NAME}/bin/eslint.js
	# create wrapper script
	@${MKDIR} ${STAGEDIR}${PREFIX}/bin
	@${ECHO_CMD} '#!/bin/sh' > ${STAGEDIR}${PREFIX}/bin/eslint
	@${ECHO_CMD} 'exec ${SETENV} NODE_PATH=${PREFIX}/lib/node_modules ${PREFIX}/lib/node_modules/${PACKAGE_NAME}/bin/eslint.js "$${@}"' >> ${STAGEDIR}${PREFIX}/bin/eslint
	@${CHMOD} +x ${STAGEDIR}${PREFIX}/bin/eslint

.include <bsd.port.mk>
