PORTNAME=	govulncheck
PORTVERSION=	1.1.4
DISTVERSIONPREFIX=	v
PORTREVISION=	2
CATEGORIES=	local security
DISTNAME=	vuln

MAINTAINER=	einar@isnic.is
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

post-patch: .PHONY
	@${REINPLACE_CMD} -e 's|%%GO_CMD%%|${GO_CMD}|g' ${WRKSRC}/internal/scan/util.go
	@${REINPLACE_CMD} -e 's|%%GO_CMD%%|${GO_CMD}|g' ${WRKSRC}/all_test.go
	@${REINPLACE_CMD} -e 's|%%GO_CMD%%|${GO_CMD}|g' ${WRKSRC}/internal/scan/run.go
	@${REINPLACE_CMD} -e 's|%%GO_CMD%%|${GO_CMD}|g' ${WRKSRC}/internal/test/packages.go
	@${REINPLACE_CMD} -e 's|%%GO_CMD%%|${GO_CMD}|g' ${WRKSRC}/internal/test/testenv.go
	@${REINPLACE_CMD} -e 's|%%GO_CMD%%|${GO_CMD}|g' ${WRKSRC}/internal/testenv/testenv.go
	@${REINPLACE_CMD} -e 's|%%GO_CMD%%|${GO_CMD}|g' ${WRKSRC}/internal/vulncheck/packages.go
	@${REINPLACE_CMD} -e 's|%%GO_CMD%%|${GO_CMD}|g' ${WRKSRC}/vendor/golang.org/x/tools/go/gcexportdata/gcexportdata.go
	@${REINPLACE_CMD} -e 's|%%GO_CMD%%|${GO_CMD}|g' ${WRKSRC}/vendor/golang.org/x/tools/go/packages/packagestest/export.go
	@${REINPLACE_CMD} -e 's|%%GO_CMD%%|${GO_CMD}|g' ${WRKSRC}/vendor/golang.org/x/tools/internal/gcimporter/exportdata.go
	@${REINPLACE_CMD} -e 's|%%GO_CMD%%|${GO_CMD}|g' ${WRKSRC}/vendor/golang.org/x/tools/internal/gocommand/invoke.go
	@${REINPLACE_CMD} -e 's|%%GO_CMD%%|${GO_CMD}|g' ${WRKSRC}/vendor/golang.org/x/tools/internal/goroot/importcfg.go
	@${REINPLACE_CMD} -e 's|%%GO_CMD%%|${GO_CMD}|g' ${WRKSRC}/vendor/golang.org/x/tools/internal/testenv/testenv.go

#internal/scan/run.go
#buf.WriteString("v0.0.0")
#post-extract: .PHONY
pre-deps post-deps-all pre-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

GH_TUPLE:=	${GH_TUPLE:Ngolang\:telemetry\:*}

.include <bsd.port.mk>
