# frozen_string_literal: true

class ApplicationSerializer
  include ActiveModel::Serialization

  delegate_missing_to :@object

  def initialize(object)
    @object = object
  end

  def as_json(*)
    serializable_hash(only: [])
  end
end
