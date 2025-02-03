PORTNAME=	ssh-manager
DISTVERSIONPREFIX=	v
DISTVERSION=	1.39.0
CATEGORIES=	local security

MAINTAINER=	none@freebsd.org
COMMENT=	SSH Key Manager for 1Password, Bitwarden and AWS S3
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	omegion

PLIST_FILES=	bin/${PORTNAME}

# looks like -buildid=${PORTVERSION} is not working
GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static"

post-patch: .PHONY
	${REINPLACE_CMD} \
		's/log\.Infoln(version)/version.Version = "${DISTVERSIONPREFIX}${PORTVERSION}"; version.Time = "${%FT%TZ:L:gmtime}"; &/' \
	${WRKSRC}/cmd/version.go

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
