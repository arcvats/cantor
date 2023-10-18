# frozen_string_literal: true

module Structure
  module List
    # DoublyNode class
    class DoublyNode
      def initialize(value = nil)
        @value = value
        @next = nil
        @previous = nil
      end
    end

    # Doubly linked list class
    class Doubly
      attr_reader :size

      def initialize
        @head = nil
        @tail = nil
        @size = 0
      end

      def empty?
        @size.zero?
      end

      def append(value = nil)
        new_node = DoublyNode.new(value)
        if empty?
          @head = new_node
        else
          new_node.previous = @tail
          @tail&.next = new_node
        end
        @tail = new_node
        @size += 1
      end

      def prepend(value = nil)
        new_node = DoublyNode.new(value)
        if empty?
          @tail = new_node
        else
          @head&.previous = new_node
          new_node.next = @head
        end
        @head = new_node
        @size += 1
      end

      def destroy
        @head = nil
        @tail = nil
        @size = 0
      end

      def delete_first
        return nil if empty?

        value = @head&.value
        if @size == 1
          destroy
          return value
        end

        @head = @head&.next
        @head&.previous = nil
        @size -= 1
        value
      end

      def delete_last
        return nil if empty?

        value = @tail&.value
        if @size == 1
          destroy
          return value
        end

        @tail = @tail&.previous
        @tail&.next = nil
        @size -= 1
        value
      end
    end
  end
end
