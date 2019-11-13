class AddForeignKeyOnCategoryIdInItems < ActiveRecord::Migration[6.0]
  def change
    add_foreign_key :items, :categories
  end
end
