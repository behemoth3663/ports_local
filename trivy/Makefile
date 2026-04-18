PORTNAME=	trivy
PORTVERSION=	0.70.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local security
PKGNAMESUFFIX=	-local

MAINTAINER=	mfechner@FreeBSD.org
COMMENT=	Security scanner written in go
WWW=		https://${GO_MODULE}

LICENSE=	APACHE20

BROKEN_i386=	not supported, see https://github.com/aquasecurity/trivy/pull/9102

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	aquasecurity
GO_MOD_DIST=	github

GO_MODULE=	github.com/${GH_ACCOUNT}/${GH_PROJECT}
GO_TARGET=	./cmd/${PORTNAME}

#CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X '${GO_MODULE}/pkg/version/app.ver=${PORTVERSION}'"
GO_ENV+=	GOEXPERIMENT=jsonv2,norandomizedheapbase64

PLIST_FILES=	bin/${PORTNAME}

post-install: .PHONY
	${MKDIR} ${STAGEDIR}${DATADIR}
	${INSTALL_DATA} ${WRKSRC}/contrib/*.tpl ${STAGEDIR}${DATADIR}/
	${FIND} -s ${STAGEDIR}${DATADIR} -not -type d | ${SORT} | \
		${SED} -e 's#^${STAGEDIR}${PREFIX}/##' | \
		${SED} -E -e '/sample$$/ s#^#@sample #; \
		s#${DATADIR_REL}/bin#@(,,555) ${DATADIR_REL}/bin#; ' >> ${TMPPLIST}
	${CAT} ${TMPPLIST}

.include <bsd.port.mk>
