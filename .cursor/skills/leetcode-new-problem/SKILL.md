---
name: leetcode-new-problem
description: >-
  Scaffolds a new LeetCode Go problem from a leetcode.com URL or pasted statement:
  creates the problem folder, solution stub with a top-of-file problem description,
  and testify table-driven tests. Never writes the algorithm solution.
  Use when the user pastes a LeetCode link, asks to start/setup a new problem,
  or says «новая задача», «создай папку», «scaffold».
---

# LeetCode: новая задача (Go)

## Goal

По ссылке (или вставленному условию) создать готовую к решению папку: описание сверху, stub функции, тесты. **Алгоритм не писать** (см. правило learning coach).

## Workflow

Скопируй чеклист и отмечай шаги:

```
- [ ] 1. Разобрать URL / условие
- [ ] 2. Создать папку и stub .go
- [ ] 3. Написать *_test.go
- [ ] 4. Проверить, что пакет компилируется
- [ ] 5. Кратко сказать, как запускать тесты
```

### 1. Разобрать URL / условие

Из URL вида `https://leetcode.com/problems/<slug>/` взять `<slug>`.

Папка и файлы:

```
<slug>/
  <slug>.go
  <slug>_test.go
```

Package name: валидный Go identifier из slug без дефисов (например `remove-duplicates-from-sorted-array-ii` → `removeduplicatesii`). Не использовать `main`.

Достать условие:

1. `WebFetch` / поиск по странице задачи.
2. Если Cloudflare/пусто — попросить пользователя вставить Statement + Examples + Constraints (или Go Code Stub с LeetCode).
3. Сигнатуру функции брать из **Go** stub LeetCode; если stub нет — вывести из примеров и уточнить у пользователя при неоднозначности.

### 2. Stub `.go`

В начале файла — блочный комментарий с описанием (русский или английский как в источнике; можно кратко, но с примерами и constraints). Затем package, stub функции **без решения**.

Шаблон:

```go
/*
<title> (LeetCode <n>)

<link>

<краткое условие>

Example 1:
  Input: ...
  Output: ...

Example 2:
  ...

Constraints:
  ...
*/
package <pkg>

func <FuncName>(...) <ReturnType> {
	// TODO: implement
	<zero return so it compiles>
}
```

Zero-return: `return 0` / `return nil` / `return false` / typed zero — что нужно для компиляции. Не писать рабочую логику.

Если папка уже существует — не перезаписывать решение пользователя; обновить/добавить только то, о чём просят (часто только тесты или описание).

### 3. Тесты `*_test.go`

Обязательно:

- `github.com/stretchr/testify/require`
- table-driven + `t.Run`
- все официальные Examples
- разумные edge cases **внутри Constraints** (не добавлять `nums == nil`/`len==0`, если в constraints `length >= 1`)

Стиль проверки — как у judge LeetCode:

- in-place / «вернуть k + префикс»: `require.Equal(t, len(expected), k)` и поэлементно `nums[i]`
- обычный return value: `require.Equal(t, want, got)`
- если функция мутирует вход — копировать слайс через `append([]T(nil), tt.in...)` до вызова

Не импортировать лишнее. Не дублировать solution-код в тестах (никаких oracle-реализаций задачи).

### 4. Компиляция

Из корня модуля:

```bash
go test ./<slug>/ -c -o /dev/null
```

(только проверка сборки; полный прогон через gotestsum — на стороне пользователя).

### 5. Ответ пользователю

Кратко на русском:

1. Что создано (пути файлов).
2. Как гонять тесты: workspace task **Tasks: Run Test Task** (`gotestsum: current package`), открыв файл в папке задачи.
3. Не писать решение и не спойлерить алгоритм — максимум намёк «с чего начать», если просят.

## Do / Don't

| Do | Don't |
|----|--------|
| Описание сверху в `.go` | Писать тело решения |
| Stub + compile-ready return | `panic` вместо return без нужды |
| testify/require table tests | `testing` only asserts / `assert` package |
| Examples + constraint-valid edges | Кейсы вне constraints |
| Указать Run Test Task | Каждый раз длинная CLI-команда без причины |

## Examples

**User:** `https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/`

**Agent:** создаёт `remove-duplicates-from-sorted-array-ii/…`, stub `removeDuplicates`, тесты по Example 1/2 + edges, без алгоритма.

**User:** ссылка + «только тесты»

**Agent:** не трогает тело функции; пишет/чинит только `*_test.go`.
