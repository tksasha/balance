# frozen_string_literal: true

RSpec.describe ActsAsHasFormula do
  let(:klass) do
    Class.new do
      include ActiveModel::Validations
      include ActiveModel::Validations::Callbacks
      include ActsAsHasFormula

      attr_accessor :sum, :formula

      def initialize(formula: nil)
        @formula = formula
      end
    end
  end

  subject { klass.new formula: '2+2' }

  describe '#calculate_formula' do
    before { subject.send :calculate_formula }

    its(:sum) { should eq 4.0 }
  end

  context 'run callback `#calculate_formula` before validation' do
    before { expect(subject).to receive(:calculate_formula) }

    it { expect { subject.valid? }.to_not raise_error }
  end
end
