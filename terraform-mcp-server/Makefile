PORTNAME=	terraform-mcp-server
PORTVERSION=	1.3.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local misc
PKGNAMESUFFIX=	-local

MAINTAINER=	none@FreeBSD.org
COMMENT=	Terraform MCP Server
WWW=		https://${GO_MODULE}

LICENSE=	MPL20

USES=		go:modules

GO_MODULE=	github.com/hashicorp/${PORTNAME}

GO_TARGET=	./cmd/${PORTNAME}

CGO_ENABLED=	0
GO_LDFLAGS+=	-w
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

PLIST_FILES=	bin/${PORTNAME}

pre-build: .PHONY
	${RM} -r ${WRKSRC}/vendor/github.com/opentracing
	${RM} -r ${WRKSRC}/vendor/go.opentelemetry.io

.include <bsd.port.mk>
