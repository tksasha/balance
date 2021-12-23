# frozen_string_literal: true

class TagSerializer < ApplicationSerializer
  def as_json(*)
    serializable_hash(only: %i[id name])
  end
end
