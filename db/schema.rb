# This file is auto-generated from the current state of the database. Instead
# of editing this file, please use the migrations feature of Active Record to
# incrementally modify your database, and then regenerate this schema definition.
#
# This file is the source Rails uses to define your schema when running `rails
# db:schema:load`. When creating a new database, `rails db:schema:load` tends to
# be faster and is potentially less error prone than running all of your
# migrations from scratch. Old migrations may fail to apply correctly if those
# migrations use external dependencies or application code.
#
# It's strongly recommended that you check this file into your version control system.

ActiveRecord::Schema.define(version: 2020_07_26_101849) do

  create_table "cashes", force: :cascade do |t|
    t.decimal "sum", precision: 10, scale: 2
    t.string "name"
    t.time "deleted_at"
    t.string "formula"
    t.integer "currency", default: 0
  end

  create_table "categories", force: :cascade do |t|
    t.string "name"
    t.boolean "income", default: false
    t.string "slug"
    t.boolean "visible", default: true
    t.integer "currency", default: 0
  end

  create_table "exchange_rates", force: :cascade do |t|
    t.integer "from", default: 0
    t.integer "to", default: 0
    t.decimal "rate", precision: 7, scale: 5
    t.date "date"
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
  end

  create_table "items", force: :cascade do |t|
    t.date "date"
    t.decimal "sum", precision: 10, scale: 2, null: false
    t.integer "category_id"
    t.string "description"
    t.datetime "created_at"
    t.datetime "updated_at"
    t.text "formula"
    t.time "deleted_at"
    t.integer "currency", default: 0
    t.index ["date", "category_id"], name: "index_balans_items_on_date_and_category_id"
    t.index ["date"], name: "index_balans_items_on_date"
  end

  create_table "user_identities", force: :cascade do |t|
    t.integer "user_id"
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
  end

  create_table "users", force: :cascade do |t|
    t.string "email"
    t.string "name"
    t.datetime "created_at", precision: 6, null: false
    t.datetime "updated_at", precision: 6, null: false
    t.string "password_digest"
  end

  create_table "versions", force: :cascade do |t|
    t.string "item_type", null: false
    t.integer "item_id", limit: 8, null: false
    t.string "event", null: false
    t.string "whodunnit"
    t.text "object", limit: 1073741823
    t.datetime "created_at"
    t.text "object_changes", limit: 1073741823
    t.index ["item_type", "item_id"], name: "index_versions_on_item_type_and_item_id"
  end

  add_foreign_key "items", "categories"
end
