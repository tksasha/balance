# frozen_string_literal: true

class Result
  attr_reader :object

  def initialize(object)
    @object = object
  end

  def success?
    is_a?(Success)
  end

  def failure?
    is_a?(Failure)
  end
end
