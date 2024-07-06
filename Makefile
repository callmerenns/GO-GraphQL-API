# Setup Variable
GO = go
GET = get
RUN = run
URL = github.com/99designs/gqlgen
URL1 = github.com/99designs/gqlgen/codegen/config@v0.17.45
URL2 = github.com/99designs/gqlgen/internal/imports@v0.17.45
URL3 = github.com/99designs/gqlgen/codegen@v0.17.45
URL4 = github.com/99designs/gqlgen@v0.17.45

# Setup Rules Default
all: download generates

# Setup Rules to Download Module
download:
	@echo "Running download"
	@echo "This may take a moment..."
	@echo ""
	$(GO) $(GET) $(URL1)
	$(GO) $(GET) $(URL2)
	$(GO) $(GET) $(URL3)
	$(GO) $(GET) $(URL4)
	@echo ""
	@echo "Download finish!"

# Setup Rules to Generate Schema
generates: download
	@echo "Running generate schema"
	@echo "This may take a moment..."
	@echo ""
	$(GO) $(RUN) $(URL) generate
	@echo ""
	@echo "Generate finish!"