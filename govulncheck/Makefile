PORTNAME=	govulncheck
PORTVERSION=	1.1.4
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local net
DISTNAME=	vuln

MAINTAINER=	none@freebsd.org
COMMENT=	Database client and tools for Go vulnerability database
WWW=		https://pkg.go.dev/golang.org/x/vuln/cmd/govulncheck

LICENSE=	BSD3CLAUSE
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	golang
GH_PROJECT=	vuln

PLIST_FILES=	bin/${PORTNAME}

GO_TARGET=	./cmd/${PORTNAME}
GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static"

#internal/scan/run.go
#buf.WriteString("v0.0.0")
#post-extract: .PHONY
pre-deps post-deps-all pre-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

GH_TUPLE:=	${GH_TUPLE:Ngolang\:telemetry\:*}

.include <bsd.port.mk>
