package cmd

import (
	"bingo-example/bootstrap"
	"bingo-example/constants"
	"bingo-example/domain/entity/book"
	"bingo-example/pkg/console"
	"bingo-example/pkg/database"
	"bingo-example/pkg/es"
	"context"
	"fmt"
	"strconv"
	"sync"

	"github.com/olivere/elastic/v7"
	"github.com/spf13/cobra"
)

// CmdImportBook 将数据库的📚信息导入到es当中
var CmdImportBook = &cobra.Command{
	Use:   "import_book",
	Short: "import book data in es",
	Run:   runImportBook,
	Args:  cobra.ExactArgs(0), // 只允许且必须传 0 个参数
}

// todo fix导入
func runImportBook(cmd *cobra.Command, args []string) {
	bootstrap.SetupDB()
	bootstrap.SetupElastic()

	page, pageSize := 1, 1000
	ctx := context.Background()
	wg := sync.WaitGroup{}

	for {
		// 从mysql获取数据
		books := book.Books{}

		err := database.DB().Model(&book.Book{}).Order("id desc").Limit(pageSize).Offset((page - 1) * pageSize).Find(&books).Error
		if err != nil || len(books) == 0 {
			break
		}

		wg.Add(1)
		go func(p int) {
			defer wg.Done()

			// 导入到es
			bulk := es.ES().Bulk()
			for _, b := range books {
				req := elastic.NewBulkIndexRequest()
				req.Index(constants.BookIndex).Id(strconv.Itoa(b.ID)).Doc(b)
				bulk.Add(req)
			}

			if rep, err := bulk.Do(ctx); err != nil {
				console.Error(fmt.Sprintf("[import book failed]%s", err.Error()))
			} else {
				console.Success(fmt.Sprintf("%d page imported %d", p, len(rep.Items)))
			}
		}(page)

		page++
	}

	wg.Wait()

	console.Success("import book succeed!")
}
