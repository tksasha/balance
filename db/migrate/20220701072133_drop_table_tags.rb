class DropTableTags < ActiveRecord::Migration[7.0]
  def change
    drop_table(:tags) if table_exists?(:tags)

    drop_table(:items_tags) if table_exists?(:items_tags)
  end
end
