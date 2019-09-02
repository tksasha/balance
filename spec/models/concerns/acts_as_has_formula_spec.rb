# frozen_string_literal: true

RSpec.describe ActsAsHasFormula do
  let(:klass) do
    Class.new do
      include ActiveModel::Validations
      include ActiveModel::Validations::Callbacks

      prepend ActsAsHasFormula

      attr_accessor :sum, :formula
    end
  end

  subject { klass.new }

  describe '#formula=' do
    before { subject.formula = '2+2' }

    its(:sum) { should eq 4.0 }

    its(:formula) { should eq '2+2' }
  end
end
