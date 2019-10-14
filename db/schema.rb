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

ActiveRecord::Schema.define(version: 2019_10_14_074747) do

  create_table "cashes", force: :cascade do |t|
    t.decimal "sum", precision: 10, scale: 2
    t.string "name"
    t.time "deleted_at"
    t.string "formula"
  end

  create_table "categories", force: :cascade do |t|
    t.string "name"
    t.boolean "income", default: false
    t.string "slug"
    t.boolean "visible", default: true
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
    t.index ["date", "category_id"], name: "index_balans_items_on_date_and_category_id"
    t.index ["date"], name: "index_balans_items_on_date"
  end

end
