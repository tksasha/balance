class RemoveDeletedAtFromCategories < ActiveRecord::Migration[5.1]
  def change
    if column_exists? :categories, :deleted_at
      remove_column :categories, :deleted_at
    end
  end
end
