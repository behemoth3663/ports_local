PORTNAME=	trivy
PORTVERSION=	0.68.1
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local security

MAINTAINER=	mfechner@FreeBSD.org
COMMENT=	Security scanner written in go
WWW=		https://${GO_MODULE}

LICENSE=	APACHE20

USES=		go:modules

GO_MODULE=	github.com/aquasecurity/${PORTNAME}
GO_TARGET=	./cmd/${PORTNAME}

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static \
	-X '${GO_MODULE}/pkg/version/app.ver=${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

GO_ENV+=	GOEXPERIMENT=jsonv2

post-install: .PHONY
	${MKDIR} ${STAGEDIR}${DATADIR}
	${INSTALL_DATA} ${WRKSRC}/contrib/*.tpl ${STAGEDIR}${DATADIR}/
	${FIND} -s ${STAGEDIR}${DATADIR} -not -type d | ${SORT} | \
		${SED} -e 's#^${STAGEDIR}${PREFIX}/##' | \
		${SED} -E -e '/sample$$/ s#^#@sample #; \
		s#${DATADIR_REL}/bin#@(,,555) ${DATADIR_REL}/bin#; ' >> ${TMPPLIST}
	${CAT} ${TMPPLIST}

.include <bsd.port.mk>
