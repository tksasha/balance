# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# This file is the source Rails uses to define your schema when running `bin/rails
# db:schema:load`. When creating a new database, `bin/rails db:schema:load` tends to
# be faster and is potentially less error prone than running all of your
# migrations from scratch. Old migrations may fail to apply correctly if those
# migrations use external dependencies or application code.
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema[7.0].define(version: 2023_02_20_192035) do
  create_table "cashes", force: :cascade do |t|
    t.decimal "sum", precision: 10, scale: 2
    t.string "name"
    t.time "deleted_at"
    t.string "formula"
    t.integer "currency", default: 0
    t.integer "supercategory", default: 1, null: false
    t.index ["name", "currency"], name: "index_cashes_on_name_and_currency", unique: true
  end

  create_table "categories", force: :cascade do |t|
    t.string "name"
    t.boolean "income", default: false
    t.boolean "visible", default: true
    t.integer "currency", default: 0
    t.integer "supercategory", default: 1, null: false
    t.index ["name", "currency"], name: "index_categories_on_name_and_currency", unique: true
  end

  create_table "items", force: :cascade do |t|
    t.date "date"
    t.decimal "sum", precision: 10, scale: 2, null: false
    t.integer "category_id"
    t.string "description"
    t.datetime "created_at", precision: nil
    t.datetime "updated_at", precision: nil
    t.text "formula"
    t.time "deleted_at"
    t.integer "currency", default: 0
    t.index ["date", "category_id"], name: "index_balans_items_on_date_and_category_id"
    t.index ["date"], name: "index_balans_items_on_date"
  end

  create_table "versions", force: :cascade do |t|
    t.string "item_type", null: false
    t.integer "item_id", limit: 8, null: false
    t.string "event", null: false
    t.string "whodunnit"
    t.text "object", limit: 1073741823
    t.datetime "created_at", precision: nil
    t.text "object_changes", limit: 1073741823
    t.index ["item_type", "item_id"], name: "index_versions_on_item_type_and_item_id"
  end

  add_foreign_key "items", "categories"
end
