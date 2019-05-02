class Consolidation < Item
  delegate :name, :slug, :income?, to: :category
end
