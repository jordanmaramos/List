package test

import (
	"list"
	"testing"
)

func TestAdd(t *testing.T) {
	// Cria uma instância de ArrayList para testar
	arrayList := list.ArrayList{}
	arrayList.Init()

	// Adiciona um valor e verifica se o último índice foi atualizado corretamente
	arrayList.Add(1)
	if arrayList.LastIndex != 1 {
		t.Errorf("O último índice deveria ser 1, mas é %d", arrayList.LastIndex)
	}

	// Adiciona mais valores até que o tamanho do array seja expandido
	for i := 2; i <= 10; i++ {
		arrayList.Add(i)
	}
	if len(arrayList.Values) != 10 {
		t.Errorf("O tamanho do array deveria ser 10, mas é %d", len(arrayList.Values))
	}

	// Adiciona um valor quando o array está cheio e verifica se o tamanho foi dobrado
	arrayList.Add(11)
	if len(arrayList.Values) != 20 {
		t.Errorf("O tamanho do array deveria ser 20, mas é %d", len(arrayList.Values))
	}
}

func TestInit(t *testing.T) {
	// Cria uma instância de ArrayList e chama Init
	arrayList := list.ArrayList{}
	arrayList.Init()

	// Verifica se os valores foram inicializados corretamente
	if len(arrayList.Values) != 5 {
		t.Errorf("O tamanho do array deveria ser 5, mas é %d", len(arrayList.Values))
	}
	if arrayList.LastIndex != 0 {
		t.Errorf("O último índice deveria ser 0, mas é %d", arrayList.LastIndex)
	}

	// Altera os valores e chama Init novamente para verificar se foram resetados
	arrayList.Values[0] = 1
	arrayList.LastIndex = 1
	arrayList.Init()
	if len(arrayList.Values) != 5 {
		t.Errorf("O tamanho do array deveria ser 5, mas é %d", len(arrayList.Values))
	}
	if arrayList.LastIndex != 0 {
		t.Errorf("O último índice deveria ser 0, mas é %d", arrayList.LastIndex)
	}
}
