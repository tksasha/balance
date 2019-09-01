# frozen_string_literal: true

class Consolidation < Item
  delegate :name, :slug, :income?, to: :category
end
