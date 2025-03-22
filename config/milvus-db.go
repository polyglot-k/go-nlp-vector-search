package config

import (
	"context"
	"fmt"
	"log"

	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"github.com/milvus-io/milvus-sdk-go/v2/entity"
)

var ctx context.Context

func init() {
	ctx = context.Background()
}

func GetMilvusDBConfig() {
	client, err := client.NewClient(ctx, client.Config{
		Address: "localhost:19530",
	})
	
	fmt.Println("Schema created successfully")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// 수정된 컬렉션 스키마
	schema := &entity.Schema{
		CollectionName: "example_collection",
		Fields: []*entity.Field{
			{
				Name:       "example_id",             // PrimaryKey 필드
				DataType:   entity.FieldTypeInt64,    // PrimaryKey는 Int64로 설정
				PrimaryKey: true,                     // PrimaryKey 설정
				AutoID:     true,                     // 자동 ID 할당
				Description: "Unique ID",             // 설명 추가
			},
			{
				Name:        "example_vector",        // 벡터 필드
				DataType:    entity.FieldTypeFloatVector,
				Description: "Embedding vector",      // 설명 추가
			},
		},
	}
	fmt.Println("Schema created successfully")
	err = client.CreateCollection(ctx, schema, 2)
	if err != nil {
		log.Fatalf("failed to create collection: %v", err)
	}

	log.Println("Collection created successfully")
}
