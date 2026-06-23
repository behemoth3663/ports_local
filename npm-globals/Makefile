PORTNAME=	globals
PORTVERSION=	17.7.0
CATEGORIES=	local devel
MASTER_SITES=	LOCAL/o1vm/:npm
PKGNAMEPREFIX=	npm-
DISTFILES=	${PORTNAME}-${DISTVERSION}${EXTRACT_SUFX}

MAINTAINER=	none@freebsd.org
COMMENT=	Global identifiers from different JavaScript environments
WWW=		https://github.com/sindresorhus/globals

LICENSE=	MIT

FETCH_DEPENDS=	npm:www/npm \
		jq:textproc/jq \
		${LOCALBASE}/share/certs/ca-root-nss.crt:security/ca_root_nss

USES=		nodejs:run

GH_ACCOUNT=	sindresorhus
#GH_PROJECT=	globals

NO_BUILD=	yes
NO_ARCH=	yes

PACKAGE_NAME=	globals

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

.include <bsd.port.mk>
