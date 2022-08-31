#

PORTNAME=	tfsec
PORTVERSION=	1.27.6
DISTVERSIONPREFIX=	v
CATEGORIES=	local security

MAINTAINER=	none@freebsd.org
COMMENT=	Static analysis powered security scanner for your terraform code

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	aquasecurity

GO_TARGET=	./cmd/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X 'github.com/${GH_ACCOUNT}/${PORTNAME}/version.Version=${PORTVERSION}'"

PLIST_FILES=	bin/${PORTNAME}

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
